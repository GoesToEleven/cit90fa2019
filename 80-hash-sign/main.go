package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

const secretKey = "astronaut juggernaut jeigermeister"

func main() {
	emailWrittenToCookie := "james@uk.gov"

	s := emailWrittenToCookie + secretKey
	sum := sha256.Sum256([]byte(s))
	hash := fmt.Sprintf("%x", sum)
	valueToStoreInCookie := emailWrittenToCookie + "|" + hash

	// set cookie
	fmt.Println("value inna cookie", valueToStoreInCookie)

	// retrieve cookie
	// two possibilities ....
	// (A)
	valueFromCookie := valueToStoreInCookie
	// (B)
	// valueFromCookie := "jenny@uk.gov" + "|" + hash

	fmt.Println("value from cookie", valueFromCookie)


	xs := strings.Split(valueFromCookie, "|")
	emailFromCookie := xs[0]
	hashCodeFromCookie := xs[1]

	s2 := emailFromCookie + secretKey
	sum2 := sha256.Sum256([]byte(s2))
	hash2 := fmt.Sprintf("%x", sum2)

	if hashCodeFromCookie != hash2 {
		fmt.Println("BUSTED - hashes not match")
		fmt.Println("hash from cookie:", hashCodeFromCookie)
		fmt.Println("hash we computed:", hash2)
		os.Exit(0)
	}
	fmt.Println("the hashes matched")
	fmt.Println("hash from cookie:", hashCodeFromCookie)
	fmt.Println("hash we computed:", hash2)
}
