package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

// Not sure what to do with this yet....
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

func compress(origFile *os.File, compressionLevel level) error {
	// Create encoder object with appropriate compression level
	compressor := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	img, format, err := image.Decode(origFile)
	fmt.Printf("\nFormat found: %s\n", format)
	if err == image.ErrFormat {
		log.Printf("\n%s file format not supported.\n\n", origFile.Name())
		wg.Done()
		return nil
	} else if err != nil {
		wg.Done()
		return err
	}
	origFile.Close()

	newFile, err := os.Create(origFile.Name() + "_compressed.png")
	if err != nil {
		wg.Done()
		return err
	}

	// Compresses the image with "best compression" setting
	fmt.Println("Encoding . . .")
	err = compressor.Encode(newFile, img)
	if err != nil {
		panic(err)
	}
	newFile.Close()

	// Decrement the waitgroup counter
	wg.Done()

	return nil
}
