package main

import (
	"errors"
	"image/png"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	path := "./imgs/DarkBB8.png"

	f, err := ReadTestFile(t, path)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer f.Close()

	// Add 1 to the waitgroup so it won't panic
	wg.Add(1)
	err = compress(f)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	cf, err := os.Open(path + "_compressed.png")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer cf.Close()

	DimensionTest(f, cf)

	err = os.Remove(path + "_compressed.png")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func DimensionTest(origFile, compFile *os.File) error {
	origConf, err := png.DecodeConfig(origFile)
	if err != nil {
		return err
	}

	compConf, err := png.DecodeConfig(compFile)
	if err != nil {
		return err
	}

	if origConf.Width != compConf.Width || origConf.Height != compConf.Height {
		return errors.New("Dimensions are incorrect. compression failed")
	}

	return nil
}

func ReadTestFile(t *testing.T, path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
