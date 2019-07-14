// http3使用ServeMux为不同的url注册不同的事件处理函数

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Dollars float32

func (d Dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]Dollars

func (db database) list(writer http.ResponseWriter, request *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(writer, "%s: %s\n", item, price)
	}
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
