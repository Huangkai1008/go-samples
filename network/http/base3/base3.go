package main

import (
	"fmt"
	"go-starter/gee"
	"log"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
