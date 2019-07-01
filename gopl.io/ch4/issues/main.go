// 将符合搜索条件的issue输出为一个表格
package main

import (
	"fmt"
	"go-starter/gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

func printfIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	github.SortResult(result)
	var monthIssues, yearIssues, distantIssues []*github.Issue
	monthAgo := time.Now().AddDate(0, -1, 0)
	yearAgo := time.Now().AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(monthAgo) {
			monthIssues = append(monthIssues, item)
		} else if item.CreatedAt.After(yearAgo) {
			yearIssues = append(yearIssues, item)
		} else {
			distantIssues = append(distantIssues, item)
		}

	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("Within a month ...")
	printfIssues(monthIssues)

	fmt.Println("Within a year ...")
	printfIssues(yearIssues)

	fmt.Println("Long long ago ...")
	printfIssues(distantIssues)
}
