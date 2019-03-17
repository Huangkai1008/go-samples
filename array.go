package main

import "fmt"

func main() {
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	fmt.Println(team)

	for k, v := range team {
		fmt.Println(k, v)
	}

	var dArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			dArray[i][j] = i + j
		}
	}

	fmt.Println(dArray)
}
