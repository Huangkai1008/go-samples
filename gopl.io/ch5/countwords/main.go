package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("words: %d\nimages: %d\n", words, images)
	}

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing %s as html:%v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	unvisited := make([]*html.Node, 0)
	unvisited = append(unvisited, n)
	for len(unvisited) > 0 {
		n = unvisited[len(unvisited)-1]
		unvisited = unvisited[:len(unvisited)-1]

		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "image" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			unvisited = append(unvisited, c)
		}
	}
	return
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}
