package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Cartoon struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Day        string
}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("searcg query failed: %s", resp.Status)
	}

	var cartoon Cartoon
	if err := json.NewDecoder(resp.Body).Decode(&cartoon); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cartoon)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		go func() {
			defer wg.Done()
			fetch(url)
		}()
	}
	wg.Wait()
}
