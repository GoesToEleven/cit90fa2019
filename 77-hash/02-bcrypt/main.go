package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// simulates password created with account creation
	p := "bananabread12"
	
	// bs, the BCRYPT HASHED password, is what we store in the database 
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("GenerateFromPassword error", err)
	}

	fmt.Println(bs)
	fmt.Printf("%x\n", bs)

	fmt.Println("compare password stored with password entered")
	// simulates password entered by user to login
	p2 := "bananabread12"

	err = bcrypt.CompareHashAndPassword(bs, []byte(p2))
	if err != nil {
		log.Fatalln("passwords do not match")
	}
	fmt.Println("passwords matched")
}
