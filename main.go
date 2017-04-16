package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	imgPath = flag.String("imgPath", "./imgs", "Path to a folder")
	wg      sync.WaitGroup
)

func main() {
	flag.Parse()

	err := filepath.Walk(*imgPath, visit)
	if err != nil {
		log.Println(err)
	}

	// Wait for all compression routines to finish before exiting
	wg.Wait()
}

// Each image that is "visited" will be encoded/compressed in an attempt to save storage
func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		// Exit the function if a directory is found
		return nil
	}

	fmt.Printf("\nVisited: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	wg.Add(1)
	go compress(file)

	fmt.Printf("Finished encoding image: %s\n", f.Name())

	return nil
}
