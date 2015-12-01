/* Modify fetch to also print the HTTP status code, found in resp.status */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url1> <url2> ...\n", os.Args[0])
		os.Exit(1)
	}
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("#######################  %s  ##########################\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: Copying %v\n", err)
			os.Exit(1)
		}
	}
}
