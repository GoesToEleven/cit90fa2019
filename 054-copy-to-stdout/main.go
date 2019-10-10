package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			err = fmt.Errorf("couldn't get %s - here's the error %w", url, err)
			log.Println(err)
			os.Exit(1)
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		fmt.Println(n)
		if err != nil {
			err = fmt.Errorf("couldn't copy using io.copy - here's the error %w", err)
			log.Panic(err)
		}

		err = resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("couldn't close resp body reading %s - error %w", url, err)
			log.Panic(err)
		}
	}
}

//!-
