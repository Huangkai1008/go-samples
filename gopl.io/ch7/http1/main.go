// http1实现一个简单的http服务器
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
	for item, price := range db {
		_, _ = fmt.Fprintf(writer, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
