package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// type, number of elements in slice
	// Read doesn't dynamically allocated response size
	io.Copy(os.Stdout, resp.Body)

	/*
		func Copy(dst Writer, src Reader) (written int64, err error) {
			return copyBuffer(dst, src, nil)
		}

		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/

}

func (logWriter) Write(bs []byte) (int, error) {
	// n is number of bytes
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
