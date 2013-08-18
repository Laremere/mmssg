package main

import (
	"bytes"
	"os"
	"path/filepath"
)

var files map[string][]byte

func loadStatic() error {
	files = make(map[string][]byte)
	return filepath.Walk("static", loadStaticFile)
}

func loadStaticFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.Mode().IsRegular() {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		_, err = buf.ReadFrom(f)
		if err != nil {
			return err
		}

		files[path] = buf.Bytes()
	}

	return nil
}
