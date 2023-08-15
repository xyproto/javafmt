# javafmt

For both Java and Kotlin, **organize imports** and then format the code with `google-java-format` or `ktlint`, depending on the file extension.

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

## Issues

This utility is a bit experimental, so the correct imports may not always be used. Bug reports and pull requests are welcome!

## General info

* Version: 0.0.3
* License: BSD-3
