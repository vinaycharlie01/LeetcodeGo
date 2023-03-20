// package main

// import "fmt"

// type Article struct {
// 	Title  string
// 	Author string
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
// }

// type Book struct {
// 	Title  string
// 	Author string
// 	Pages  int
// }

// func (b Book) String() string {
// 	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
// }

// type Stringer interface {
// 	String() string
// }

// func Print(s Stringer) {
// 	fmt.Println(s.String())
// }

// func main() {
// 	a := Article{
// 		Title:  "Understanding Interfaces in Go",
// 		Author: "Sammy Shark",
// 	}
// 	Print(a)

// 	b := Book{
// 		Title:  "All About Go",
// 		Author: "Jenny Dolphin",
// 		Pages:  25,
// 	}
// 	Print(b)
// }

package main

import (
	"fmt"
)

type Working interface {
	Sound() string
}

type Person struct {
	Name  string
	Age   int
	Hight float64
}

type Aminal struct {
	Name  string
	Age   string
	Hight float64
}

func (r *Person) Sound() string {
	return r.Name
}

func (r *Aminal) Sound() string {
	return r.Name
}

func Sound(S ...Working) string {
	for _, v := range S {
		return v.Sound()
	}
	return ""
}


func passarr(a [4]int)  {
	
}

func main() {
	a:=[...]int{}

	// a := []Working{
	// 	&Person{"vinnay", 12, 5.6},
	// 	&Aminal{"dog", "2", 2.2},
	// }
	// b := Sound(a...)
	// fmt.Println(b)

	var a interface{} = "Hello, world!"
	switch a.(type) {
	case string:
		fmt.Println("this is string type")
	}


	strings.
}
