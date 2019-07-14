// httpHtml list改为输出html表格

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Dollars float32

func (d Dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]Dollars

func (db database) list(writer http.ResponseWriter, request *http.Request) {
	result := template.Must(template.ParseFiles("gopl.io/ch7/httphtml/tpl.gohtml"))
	_ = result.Execute(writer, db)
}

func (db database) price(writer http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(writer, "no such item: %q\n", item)
		return
	}
	_, _ = fmt.Fprintf(writer, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
