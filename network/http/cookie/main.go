package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	cookie1 := &http.Cookie{
		Name:  "yummy_cookie",
		Value: "choco",
	}
	cookie2 := &http.Cookie{
		Name:  "tasty_cookie",
		Value: "strawberry",
	}
	http.SetCookie(w, cookie1)
	http.SetCookie(w, cookie2)
	_, _ = fmt.Fprintln(w, "Set cookie successful")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Get cookie successful: ")
	for _, cookie := range r.Cookies() {
		_, _ = fmt.Fprintf(w, "Cookie name is %s, Cookie value is %v\n", cookie.Name, cookie.Value)
	}
}
