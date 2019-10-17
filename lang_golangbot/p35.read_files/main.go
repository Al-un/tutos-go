package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var root = "/home/al-un/git/go/src/github.com/al-un/tutos-go/lang_golangbot/p35.read_files/"

func main() {
	// in code path
	data, err := ioutil.ReadFile(root + "test.txt")
	if err != nil {
		fmt.Println("File reading error: ", err)
		return
	}
	fmt.Println("Contents of file: ", string(data))
	fmt.Println("----------------------------")

	// argument
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("Value of fpath is ", *fptr)
	data2, err2 := ioutil.ReadFile(*fptr)
	if err2 != nil {
		fmt.Println("File reading error ", err)
		return
	}
	fmt.Println("Contents of file: ", string(data2))
	fmt.Println("----------------------------")

	// bundling
	// box := packr.NewBox("../p36.read_files")
	// data3 := box.String("test.txt")
	// fmt.Println("Contents of file: ", data3)

	// chunk of file, from "Argument"
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error when reading file: ", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
	fmt.Println("----------------------------")

	// read line by line
	f, err = os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------------------------")
}
