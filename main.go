package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/zserge/webview"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

const (
	windowWidth  = 800
	windowHeight = 900
)

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func main() {
	port := flag.String("p", "9999", "port to serve on")
	directory := flag.String("d", "./build", "the directory of static file to host")
	flag.Parse()

	go func() {
		fileServer := http.FileServer(FileSystem{http.Dir(*directory)})
		http.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fileServer))
		log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}()

	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "Simple",
		Resizable: true,
		URL:       "http://127.0.0.1:9999",
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()
}
