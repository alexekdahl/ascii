# ASCII Art Generator in Go

![Go Version](https://img.shields.io/badge/go-v1.21-blue)
![License](https://img.shields.io/badge/license-MIT-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexekdahl/ascii)](https://goreportcard.com/report/github.com/alexekdahl/ascii)
![Tests](https://img.shields.io/badge/tests-100%25-success)
![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen)

## Overview

`ascii` is a Go library designed to convert images into ASCII art. The library provides a simple and efficient way to generate ASCII representations of images with color support.

## Features

- Image Decoding: Supports multiple image formats including PNG, JPEG, and GIF.
- Byte Array Support: Operates directly on byte arrays for image manipulation.
- Image Scaling: Efficiently scales images to fit the ASCII grid.
- Color Support: Outputs colored ASCII art based on the original image colors.

## Installation

Install the package using `go get`:

```bash
go get github.com/alexekdahl/ascii
```

## Usage

Here's a simple example to generate ASCII art from an image:

```go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alexekdahl/ascii"
)

func main() {
	resp, err := http.Get("https://avatars.githubusercontent.com/u/1024025?v=4")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	customASCII, err := ascii.ToASCII(body, ascii.WithChar('█'))
	if err != nil {
		panic(err)
	}

	defaultASCII, err := ascii.ToASCII(body)
	if err != nil {
		panic(err)
	}

	fmt.Println(customASCII)
	fmt.Println(defaultASCII)
}
```


## Contributing

We welcome contributions!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
