package main

import (
	"fmt"
)

var preReqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	var index int
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				index++
				order[index] = item
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(preReqs) {
		fmt.Printf("%d:\t%s\n", i, course)
	}
}
