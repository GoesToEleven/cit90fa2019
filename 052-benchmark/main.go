package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s1 := echo1()
	s3 := echo3()
	fmt.Println(s1)
	fmt.Println(s3)
}

func echo1() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " \n"
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[:], " \n")
}
