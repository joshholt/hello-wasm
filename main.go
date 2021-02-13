//go:generate broccoli -src web
// +build !wasm

package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	http.Handle("/foo", br.Serve("web"))
	http.Handle("/", &app.Handler{
		Title:       "Hello WASM",
		Name:        "Hello WASM",
		Description: "An example go-app based WASM PWA",
		Styles: []string{
			"/web/css/hello.css",
		},
		Resources: BroccoliDir("web"),
	})

	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
