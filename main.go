package main

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var err error

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Not enough arguments")
	}

	// Set quality for JPEG compression
	quality := 100
	if len(args) == 2 {
		quality, err = strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
	}

	filename := args[0]
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	ext := filepath.Ext(filename)
	ext = strings.ToLower(ext)
	name := filepath.Base(filename)
	name = name[:len(name)-len(ext)]

	// Convert image
	if ext == ".png" {
		err = ConvertPngToJpeg(input, name, quality)
		if err != nil {
			panic(err)
		}
	} else if ext == ".jpeg" || ext == ".jpg" {
		err = ConvertJpegToPgn(input, name)
		if err != nil {
			panic(err)
		}
	}

}

func ConvertPngToJpeg(input *os.File, name string, quality int) error {
	image, err := png.Decode(input)
	if err != nil {
		return err
	}

	output, err := os.Create(name + ".jpeg")
	if err != nil {
		return err
	}

	options := new(jpeg.Options)
	options.Quality = quality

	err = jpeg.Encode(output, image, options)
	if err != nil {
		return err
	}

	return nil
}

func ConvertJpegToPgn(input *os.File, name string) error {
	image, err := jpeg.Decode(input)
	if err != nil {
		return err
	}
	output, err := os.Create(name + ".png")
	if err != nil {
		return err
	}

	err = png.Encode(output, image)
	if err != nil {
		return err
	}

	return nil
}

func CompressJpeg() {
}
