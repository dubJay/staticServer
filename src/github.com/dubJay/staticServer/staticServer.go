package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var (
	port      = flag.String("port", ":8080", "Port for server to listen on")
	rootDir   = flag.String("rootDir", "", "Path to webdir structure")
	resources = flag.String("resources", "resources", "Images directory")
	static    = flag.String("static", "static" , "CSS, JS, etc...")
)

func main() {
	flag.Parse()

	staticFS := http.FileServer(http.Dir(filepath.Join(*rootDir, *static)))
	imageFS := http.FileServer(http.Dir(filepath.Join(*rootDir, *resources)))
	http.Handle("/static/", http.StripPrefix("/static", staticFS))
	http.Handle("/images/", http.StripPrefix("/images", imageFS))

	log.Fatal(http.ListenAndServe(*port, nil))
}
