package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	imgPath = flag.String("imgPath", "./DarkBB8.png", "Path to an image")
)

func main() {
	flag.Parse()

	// Open the img file for decoding
	f, err := os.Open(*imgPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// Create the file to be written
	newF, err := os.Create("compressed.png")
	defer newF.Close()
	if err != nil {
		panic(err)
	}

	// Call the compress func (program will hang)
	compress(f, newF, High)

	fmt.Println("Finished encoding image")
}
