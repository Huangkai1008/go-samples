// http2实现根据不同url匹配的http服务器
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

func (db database) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/list":
		for item, price := range db {
			_, _ = fmt.Fprintf(writer, "%s: %s\n", item, price)
		}
	case "/price":
		item := request.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			writer.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintf(writer, "no such item: %q\n", item)
			return
		}
		_, _ = fmt.Fprintf(writer, "%s\n", price)
	default:
		msg := fmt.Sprintf("no such page: %s\n", request.URL)
		http.Error(writer, msg, http.StatusNotFound)
	}

}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
