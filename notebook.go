package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"golang.org/x/net/webdav"
)

const (
	targetDir = "webdav"
)

func main() {
	e := echo.New()

	srv := &webdav.Handler{
		FileSystem: webdav.Dir(targetDir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			log.Printf("WEBDAV: %#s, ERROR: %v", r, err)
		},
	}
	e.GET("/webdav", echo.WrapHandler(srv))
	e.Static("/", "superlight")
	e.POST("/save", func(c echo.Context) error {
		log.Println(c.FormParams())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))

}
