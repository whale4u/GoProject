package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// 组合
type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

func main() {
	author1 := author{
		"Niko",
		"Buck",
		"Make a difference!",
	}

	post1 := post{
		"Learning go",
		"go go go",
		author1,
	}

	post1.details()
}
