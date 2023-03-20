package main

import "fmt"

type Len interface {
	Len() int
}

type List struct {
	Arr []int
}

func NewList(arr []int) *List {
	return &List{Arr: arr}
}

type String struct {
	Str string
}

type Map struct {
	M map[int]int
}

func (L *List) Len() int {
	count := 0
	for _, _ = range L.Arr {
		count++
	}
	return count
}

func (L *Map) Len() int {
	count := 0
	for _, _ = range L.M {
		count++
	}
	return count
}

func (S *String) Len() int {
	count := 0
	for _, _ = range S.Str {
		count++
	}
	return count
}

func Lenth(l Len) int {
	return l.Len()
}

func main() {

	// a := List{[]int{1, 2, 3, 4, 5}}
	a := NewList([]int{1, 2, 3, 4})
	fmt.Println(Lenth(a))
	b1 := String{"Vinafhfhfghghy"}
	fmt.Println(Lenth(&b1))

	c1 := Map{map[int]int{
		1: 10,
		2: 20,
		3: 30,
		4: 10,
		5: 20,
		6: 30,
	}}
	fmt.Println(Lenth(&c1))

}
