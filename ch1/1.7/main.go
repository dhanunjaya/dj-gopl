/* Exercise 1.7: The function call io.Copy(dst, src) reads from src and
 * and writes to dst. use it instead of ioutil.ReadAll to copy the
 * response body to os.Stdout without requiring a buffer large enough
 * to hold the entire stream. Be sure to check the result of io.Copy.
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url1> <url2> ...\n", os.Args[0])
		os.Exit(1)
	}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: Copying %v\n", err)
			os.Exit(1)
		}
	}
}
