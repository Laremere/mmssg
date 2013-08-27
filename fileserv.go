package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var files map[string][]byte
var fileChange map[string]time.Time

func loadStatic() error {
	files = make(map[string][]byte)
	fileChange = make(map[string]time.Time)
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

		fileID := "/" + strings.Replace(path, "\\", "/", -1)
		files[fileID] = buf.Bytes()
		fileChange[fileID] = info.ModTime()
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
		http.ServeContent(w, r, path, fileChange[path], bytes.NewReader(source))
	} else {
		log.Println("Invalid Static Request: ", r.URL.Path)
		http.NotFound(w, r)
	}
}
