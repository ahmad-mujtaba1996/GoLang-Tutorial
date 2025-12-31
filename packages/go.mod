module github.com/ahmad-mujtaba1996/GoLang-Tutorial

go 1.25.5

// Dependencies for color package. Indirect means they are not directly used in our code but are required by other packages we use.
require (
	github.com/fatih/color v1.18.0
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)

// go.sum file contains the expected cryptographic checksums of the content of specific module versions. It ensures that the modules we download have not been tampered with.
