package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"A", "B"},
	},
	{
		Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"C"},
	},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
