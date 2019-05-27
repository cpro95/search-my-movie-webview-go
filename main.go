package main

import (
	"net/http"
	"log"

	"github.com/zserge/webview"
)

func startServer() string {
	go func() {
		http.Handle("/", http.FileServer(assetFS()))
		
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	return "http://localhost:8080"
}


func main() {

	url := startServer()
	w := webview.New(webview.Settings{
		Width:  800,
		Height: 980,
		Title:  "Search My Movie",
		URL:    url,
	})
	defer w.Exit()
	w.Run()
}
