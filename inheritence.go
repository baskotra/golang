package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Name of author:", b.author.firstName+b.author.lastName)
	fmt.Println("Bio of author: ", b.bio)
	fmt.Println("Title of blog: ", b.title)
	fmt.Println("Content of blog: ", b.content)
	fmt.Println(b.fullName())
}

func (a author) fullName() string {
	return a.firstName + a.lastName
}

type website struct {
	blogPost []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of website")
	for _, v := range w.blogPost {
		v.details()
		fmt.Println()
	}
}

func main() {
	a := author{
		"David",
		"Googins",
		"NAVY SEAL",
	}
	b := blogPost{
		"Can't Hurt me",
		"Bio",
		a,
	}
	b2 := blogPost{
		"You can't see me",
		"Motivation",
		a,
	}

	w := website{
		blogPost: []blogPost{b, b2},
	}
	w.contents()
}
