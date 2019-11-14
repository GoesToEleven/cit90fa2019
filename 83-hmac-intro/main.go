package main

import (
	"crypto/hmac"
	"log"
	"crypto/sha512"
	"fmt"
)

var key = []byte("astronaut juggernaut flouride")

func main() {
	emailWritten := "mike@uk.gov"
	bs, err := signMessage([]byte(emailWritten))
	if err != nil {
		log.Fatalln("error signing Message")
	}

	// emailRetrieved := "mike@uk.gov"
	emailRetrieved := "mac@uk.gov"
	
	b, err := checkSig([]byte(emailRetrieved), bs)
	if err != nil {
		log.Fatalln("error checking sig")
	}
	fmt.Println("were the messages the same? - ", b)
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in signMessage writing to hash %w", err)
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}


