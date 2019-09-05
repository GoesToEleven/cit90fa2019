package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("war-peace.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	var count int
	m := map[string]int{}
	for s.Scan() {
		// t := s.Text()
		// fmt.Println(t)
		count++
		m[s.Text()]++
	}
	fmt.Println(count)

	for k, v := range m {
		fmt.Printf("word: %s \t\t occurs: %d \n", k, v)
	}

	// func (r receiver) identifier(p parameters) returns { code }
}
