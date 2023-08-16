# javafmt

For both Java and Kotlin, **automatically add or remove imports** and then format the code with `google-java-format` or `ktlint`, depending on the file extension.

## Requirements

* `google-java-format`
* `ktlint`

## Example use

```
cd my-java-project
javafmt
```

Or to change the files (`-w` for "write"):

```
cd my-java-project
javafmt -w
```

Or to format all `.kt` and `.java` in a directory, but not touch the imports:

```
cd my-java-project
javafm -n -w
```

## Quick installation

Requires Go 1.21 or later:

    go install github.com/xyproto/javafmt@latest

## Issues

This utility is a bit experimental, so the correct imports may not always be used. Bug reports and pull requests are welcome!

## General info

* Version: 1.0.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
