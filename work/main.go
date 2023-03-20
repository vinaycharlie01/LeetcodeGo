// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Worker(i int, recive <-chan int, send chan<- int) {
// 	for v := range recive {
// 		fmt.Println("started", "id", i, "work", v)
// 		// time.Sleep(time.Second)
// 		fmt.Println("Finish", "id", i, "work", v)
// 		send <- v * 2
// 	}
// }

// func main() {

// 	var wg sync.WaitGroup
// 	ch := make(chan int)
// 	ch2 := make(chan int)
// 	for i := 0; i <= 10; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			Worker(i, ch, ch2)
// 		}(i)
// 	}
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 		defer close(ch2)
// 	}()

// 	for i := 0; i < 3; i++ {
// 		<-ch2
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	ch := make(chan int)
// 	var wg sync.WaitGroup
// 	// var mu sync.Mutex

// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			// mu.Lock()
// 			ch <- i
// 			// mu.Unlock()
// 		}(i)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()
// 	for v := range ch {
// 		fmt.Println(v)
// 	}

// }

package main

import (
	"fmt"
	"sync"
)

var mes string

var mu sync.Mutex

func Hello(a string) {
	mu.Lock()
	mes = a
	mu.Unlock()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go Hello("vinahy")
	go Hello("fnfjfj")
	wg.Wait()

	fmt.Println(mes)

}
