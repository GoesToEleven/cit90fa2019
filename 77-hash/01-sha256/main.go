package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s1 := "here is some arbtrary text"
	s2 := "here is some arbtrary textt"

	sum1 := sha256.Sum256([]byte(s1))
	sum2 := sha256.Sum256([]byte(s2))

	fmt.Printf("sum1:\t %x\n", sum1)
	fmt.Printf("sum2:\t %x\n", sum2)
	fmt.Println("sum1==sum2", sum1 == sum2)
}

/*
WHAT THIS TELLS ME
what was intended to be sent
is what was received
*/

// their number when they put the 0's & 1's of the file through sha256
// 9f0721551a24a1eb43d2005cd58bd7b17574e50384b8da8896b0754259790752

// my number when I put the 0's & 1's of the file through sha256
// 9f0721551a24a1eb43d2005cd58bd7b17574e50384b8da8896b0754259790752
