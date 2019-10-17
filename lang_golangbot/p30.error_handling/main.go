package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	// Handle Default path error
	_, err := os.Open("/anInvalidFilePath.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path ", err.Path, " failed to open")
		// return
	}
	// fmt.Println(f.Name(), " successfully opened")

	// DNS Errors
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("Operation timed out")
		} else if err.Temporary() {
			fmt.Println("Temporary error")
		} else {
			fmt.Println("Generic error: ", err)
		}
		// return
	}
	fmt.Println("Address is ", addr)

	// Direct comparison errors
	files, error := filepath.Glob("[/home/*/al-un/git/go/src/*")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		// return
	}
	fmt.Println("Matched files: ", files)
}
