package main

import "fmt"

type grade int

type student struct {
	first, last string
	score       grade
}

type result map[string]int

func main() {
	mike := student{
		first: "Mike",
		last:  "Hallstone",
		score: 42,
	}
	fmt.Println(mike)

	r := result{
		"taqueria Z":      78,
		"ponchos eateria": 84,
	}
	fmt.Println(r)
}
