package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
)

var (
	imgPath = flag.String("imgPath", "./DarkBB8.png", "Path to an image")
)

type level int

func (l level) toInt() int {
	return int(l)
}

const (
	low    level = 1
	medium level = 5
	high   level = 10
)

func main() {
	flag.Parse()

	// Open the img file for decoding
	f, err := os.Open(*imgPath)
	if err != nil {
		panic(err)
	}

	// Create the file to be written
	newF, err := os.Create("compressed.png")
	if err != nil {
		panic(err)
	}

	// Call the compress func (program wil hang)
	compress(f, newF, high)

	fmt.Println("Finished encoding image")
}

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

	// Compresses the image with "best compression" setting
	fmt.Println("Encoding . . .")
	fmt.Printf("Compression level: %d\n", compressionLevel.toInt())
	for i := 0; i < compressionLevel.toInt(); i++ {
		err = compressor.Encode(newFile, img)
		if err != nil {
			return err
		}
		fmt.Printf("Finished compression level: %d\n", i+1)
	}

	return nil
}
