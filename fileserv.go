package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

		files["/"+strings.Replace(path, "\\", "/", -1)] = buf.Bytes()
	}

	return nil
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/static/index.html"
	}

	source, ok := files[path]
	if ok {
		_, err := w.Write(source)
		if err != nil {
			http.Error(w, "error getting file", http.StatusInternalServerError)
			log.Println("File request failed: ", err)
		}
	} else {
		log.Println("Invalid Static Request: ", r.URL.Path)
		http.NotFound(w, r)
	}
}
