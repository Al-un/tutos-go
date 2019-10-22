package beginner

import (
	"encoding/json"
	"fmt"
)

type book struct {
	Title  string `json:"title"`
	Author author `json:"author"`
}

type author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

// RunJSON https://tutorialedge.net/golang/go-json-tutorial/
func RunJSON() {
	// JSON-ing
	author := author{Sales: 3, Age: 25, Developer: true}
	book := book{Title: "Learning Concurreny in Python", Author: author}
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	// Indentation
	byteArray, err = json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

}
