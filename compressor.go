package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func compress(origFile *os.File) error {
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
