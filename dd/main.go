package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
)

type Printer interface {
	Print()
	Scan()
	Fax()
}

type OfficePrinter struct{}

func (p OfficePrinter) Print() {
	fmt.Println("Office Printer: Print")
}

func (p OfficePrinter) Scan() {
	fmt.Println("Office Printer: Scan")
}

func (p OfficePrinter) Fax() {
	fmt.Println("Office Printer: Fax")
}

type HomePrinter struct{}

func (p HomePrinter) Print() {
	fmt.Println("Home Printer: Print")
}

func (p HomePrinter) Scan() {
	fmt.Println("Home Printer: Scan")
}
func (p HomePrinter) Fax() {
	fmt.Println("Office Printer: Fax")
}

func Person1(a int) int {
	var Person2 func(int) int
	Person2 = (func(a int) int {
		var person3 func(int) int
		person3 = func(i int) int {
			return i + 10
		}
		return person3(a)
	})
	return Person2(a)
}

func ReadLine(reader *bufio.Reader) string {
	val, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(string(val), "\r")
}

func Incr(val []*int) {
	for i, j := 0, len(val)-1; i < j; i, j = i+1, j-1 {
		*val[i], *val[j] = *val[j], *val[i]
	}
	fmt.Println(val)
}

func Pointer(val *[]int) {
	sort.Slice(*val, func(i, j int) bool {
		return (*val)[i] < (*val)[j]
	})
	fmt.Println(val)
}

func Pointer2(arr []int) []int {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sort.Ints(arr)
	}()
	wg.Wait()
	return arr
}

func main() {
	// a := [...]int{2, 1, 3}
	// // a = append(a, 10)
	// Pointer((*[4]int)())

	// Pointer(10, 20, 30, 40, 50)
	// a := []*int{, 20, 30, 40}
	// /Incr(a)

	// reader := bufio.NewReader(os.Stdin)

	// val, _, err := reader.ReadLine()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// intval, err := strconv.ParseInt(strings.TrimSpace(string(val)), 2, 64)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(intval)

	// var officePrinter Printer = OfficePrinter{}
	// officePrinter.Print() // "Office Printer: Print"
	// officePrinter.Scan()  // "Office Printer: Scan"
	// officePrinter.Fax()   // "Office Printer: Fax"

	// var homePrinter Printer = HomePrinter{}
	// homePrinter.Print() // "Home Printer: Print"
	// homePrinter.Scan()  // "Home Printer: Scan"
	// homePrinter.Fax()   // error: HomePrinter does not implement Fax()
	// var a int = 10
	// func(a int) {
	// 	fmt.Println(a)
	// }(10)
	// fmt.Println(a)

	// func() {
	// 	// define a variable that is local to the IIFE
	// 	message := "Hello, world!"

	// }()
	// fmt.Println(message)

	//what is the diffrence between two  function

	// a := func() {

	// }
}
