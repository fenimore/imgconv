package main

import (
	"os"
	"path/filepath"

	"image/jpeg"
	"image/png"
)

func main() {
	filename := os.Args[1]
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	ext := filepath.Ext(filename)
	name := filepath.Base(filename)
	name = name[:len(name)-len(ext)]
	if ext == ".png" {
	}

	image, err := png.Decode(input)
	if err != nil {
		panic(err)
	}

	output, err := os.Create(name + ".jpeg")
	if err != nil {
		panic(err)
	}

	options := new(jpeg.Options)
	options.Quality = 100

	err = jpeg.Encode(output, image, options)
	if err != nil {
		panic(err)
	}
}
