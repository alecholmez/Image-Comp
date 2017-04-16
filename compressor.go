package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type level int

func (l level) toInt() int {
	return int(l)
}

const (
	// Low amount of compression
	Low level = 1
	// Medium amount of compression
	Medium level = 5
	// High amount of compression
	High level = 10
)

func compress(origFile *os.File, newFile *os.File, compressionLevel level) error {
	// Create encoder object with appropriate compression level
	compressor := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	img, format, err := image.Decode(origFile)
	fmt.Printf("Format found: %s\n", format)
	if err != nil {
		return err
	}
	origFile.Close()

	// Compresses the image with "best compression" setting
	fmt.Println("Encoding . . .")
	err = compressor.Encode(newFile, img)
	if err != nil {
		return err
	}
	newFile.Close()

	return nil
}
