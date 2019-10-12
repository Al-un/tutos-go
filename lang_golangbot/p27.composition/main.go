package main

import "fmt"

type author struct {
	first string
	last  string
	bio   string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.first, a.last)
}

type post struct {
	title   string
	content string
	author
	co author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("cO-Author: ", p.co.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Printf("Contents of website\n")
	for _, p := range w.posts {
		p.details()
		fmt.Println()
	}
}

func main() {
	a := author{
		first: "Jack",
		last:  "Dickinson",
		bio:   "some bio",
	}
	b := author{
		first: "Co-Jack",
		last:  "Co-Dickinson",
		bio:   "Co-some bio",
	}

	p := post{
		title:   "title",
		content: "content",
		author:  a,
		co:      b,
	}
	p2 := post{
		title:   "title2",
		content: "content2",
		author:  b,
		co:      a,
	}
	p3 := post{
		title:   "title3",
		content: "content3",
		author:  b,
		co:      b,
	}

	w := website{
		posts: []post{p, p2, p3},
	}
	w.contents()
}
