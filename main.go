package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/xyproto/autoimport"
)

const versionString = "javafmt 0.0.1"

func organizeImports(data []byte, onlyJava, removeExistingImports, verbose bool) []byte {
	ima, err := autoimport.New(onlyJava, removeExistingImports)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not organize imports: %v\n", err)
		return data
	}
	newData, err := ima.FixImports(data, verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not fix imports: %v\n", err)
		return data
	}
	return newData
}

func formatJava(data []byte) ([]byte, error) {
	cmd := exec.Command("google-java-format", "-a")
	cmd.Stdin = bytes.NewReader(data)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func formatKotlin(data []byte) ([]byte, error) {
	cmd := exec.Command("ktlint", "--format", "--stdin")
	cmd.Stdin = bytes.NewReader(data)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func formatData(filename string, data []byte, writeToFiles, verbose bool) {
	var err error
	var newData []byte

	ext := filepath.Ext(filename)
	switch ext {
	case ".java":
		data = organizeImports(data, true, false, verbose)
		newData, err = formatJava(data)
	case ".kt":
		data = organizeImports(data, false, false, verbose)
		newData, err = formatKotlin(data)
	default:
		return
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting %s: %v\n", filename, err)
		return
	}

	if writeToFiles {
		if err := ioutil.WriteFile(filename, newData, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", filename, err)
			return
		}
	} else {
		fmt.Print(string(newData))
	}
}

func formatFile(filename string, writeToFiles, verbose bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		return
	}
	formatData(filename, data, writeToFiles, verbose)
}

func main() {
	var showVersion bool
	var showHelp bool
	var writeToFiles bool
	var verbose bool

	pflag.BoolVarP(&showVersion, "version", "V", false, "Print version")
	pflag.BoolVar(&showHelp, "help", false, "Show help")
	pflag.BoolVarP(&writeToFiles, "write", "w", false, "Write changes back to the files instead of outputting to stdout")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")

	pflag.Usage = func() {
		fmt.Printf("Usage of %s:\n\n", os.Args[0])
		fmt.Println("A utility for formatting Java and Kotlin code.")
		fmt.Println("It organizes imports and formats the code using `google-java-format` for Java files and `ktlint` for Kotlin files.\n")
		pflag.PrintDefaults()
		fmt.Println()
	}

	pflag.Parse()

	if showHelp {
		pflag.Usage()
		return
	}

	if showVersion {
		fmt.Println(versionString)
		return
	}

	args := flag.Args()
	if len(args) == 0 && !isStdinAvailable() {
		// No argument and not reading from stdin; process all *.java and *.kt files in the current directory
		files, _ := filepath.Glob("*.java")
		kotlinFiles, _ := filepath.Glob("*.kt")
		files = append(files, kotlinFiles...)
		for _, f := range files {
			formatFile(f, writeToFiles, verbose)
		}
	} else if len(args) == 0 && isStdinAvailable() {
		// Reading from stdin
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
		formatData("stdin", data, false, verbose)
	} else {
		// Format specific files or patterns
		for _, pattern := range args {
			files, _ := filepath.Glob(pattern)
			for _, f := range files {
				formatFile(f, writeToFiles, verbose)
			}
		}
	}
}

func isStdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
