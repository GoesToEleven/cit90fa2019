package main

import "fmt"

type listItem struct {
	value    int
	up, down *listItem
}

func main() {
	first := listItem{
		value: 1,
	}
	second := listItem{
		value: 2,
	}
	third := listItem{
		value: 3,
	}

	first.down = &second
	second.up = &first
	second.down = &third
	third.up = &second

	fmt.Println("first", first)
	fmt.Println("first", second)
	fmt.Println("first", third)
}
