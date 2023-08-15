# javafmt

For both Java and Kotlin, organize imports and then format the code with either `google-java-format` or `ktlint`.

NOTE: This utility is a bit experimental and has only been tested on Java code so far.

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

## Quick installation

    go install github.com/xyproto/javafmt@latest

## General info

* Version: 0.0.3
* License: BSD-3
