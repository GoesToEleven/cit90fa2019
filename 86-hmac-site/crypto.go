package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

// secretKey DO NOT SHARE
const secretKey = "home florpus magic astronaut"

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, []byte(secretKey))
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
