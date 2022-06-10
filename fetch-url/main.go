// It fetches and print the content in a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading %s: %v", url, err)
			os.Exit(1)
		}
	}
	fmt.Printf("\nTime elapsed: %vs\n", time.Since(start).Seconds())
}
