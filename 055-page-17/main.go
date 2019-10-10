package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		switch {
		// www.google.com
		case !strings.HasPrefix(url, "http"):
			url = "https://" + url
		// http://www.google.com
		case strings.HasPrefix(url, "http"):
			url = "https" + strings.TrimPrefix(url, "http")
		}

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

		for k, v := range resp.Header {
			fmt.Println(k, v)
		}
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Status)
		fmt.Println(resp.Request.URL)

		err = resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("couldn't close resp body reading %s - error %w", url, err)
			log.Panic(err)
		}
	}
}

//!-
