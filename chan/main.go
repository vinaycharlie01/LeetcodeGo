package main

import (
	"fmt"
	"sync"
)

func Sending(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
	}
}

func REciver(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func Mergech(left chan int, right chan int, c chan int) {
	// defer close(c)
	val, ok := <-left
	val2, ok2 := <-right
	for ok && ok2 {
		if val < val2 {
			c <- val
			val, ok = <-left
		} else {
			c <- val2
			val2, ok2 = <-right
		}
	}
	for ok {
		c <- val
		val, ok = <-left
	}
	for ok2 {
		c <- val2
		val2, ok2 = <-right
	}

}
func MergeSort(arr []int, ch chan int) {
	if len(arr) < 2 {
		ch <- arr[0]
		defer close(ch)
		return
	}
	left := make(chan int)
	right := make(chan int)
	go MergeSort(arr[len(arr)/2:], left)
	go MergeSort(arr[:len(arr)/2], right)
	go Mergech(left, right, ch)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3}
	var wg sync.WaitGroup
	ch := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < len(a); i++ {
			ch <- a[i]
		}
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		for i := 0; i < len(b); i++ {
			ch2 <- b[i]
		}
	}()
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(c)
		Mergech(ch, ch2, c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	wg.Wait()
}
