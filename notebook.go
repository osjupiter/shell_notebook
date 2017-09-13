package main

import (
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

const (
	targetDir = "webdav"
)

func main() {

	srv := &webdav.Handler{
		FileSystem: webdav.Dir(targetDir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			log.Printf("WEBDAV: %#s, ERROR: %v", r, err)
		},
	}
	http.Handle("/", srv)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error with WebDAV server: %v", err)
	}

}
