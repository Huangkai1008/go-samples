package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyForm(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == " " {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if len(args) < 3 || os.Args[2] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	fmt.Println(s)
	return s
}
