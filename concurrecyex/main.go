// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// func Proccessdata(a string, ch chan string) {
// 	defer close(ch)
// 	res, err := http.Get(a)
// 	if err != nil {
// 		ch <- fmt.Sprintf("%s is down\n", a)
// 		return
// 	}
// 	defer res.Body.Close()
// 	ch <- fmt.Sprintf("%s is-->%v\n", a, res.Status)
// 	time.Sleep(time.Second * 2)
// }

// func main() {
// 	// urls := []string{
// 	// 	"https://www.google.com",
// 	// 	"https://www.yahoo.com",
// 	// 	"https://www.amazon.com",
// 	// 	"https://www.stackoverflow.com",
// 	// 	"https://www.gitfnfhhub.com",
// 	// }

// 	// ch := make(chan string)
// 	// var wg sync.WaitGroup

// 	// for i, url := range urls {
// 	// 	wg.Add(1)
// 	// 	go func(url string, i int) {
// 	// 		defer wg.Done()
// 	// 		resp, err := http.Get(url)
// 	// 		if err != nil {
// 	// 			ch <- fmt.Sprintf("%s is down\n and %v", url, i)
// 	// 			return
// 	// 		}
// 	// 		defer resp.Body.Close()
// 	// 		ch <- fmt.Sprintf("%s -> %v\n and %v", url, resp.Status, i)
// 	// 	}(url, i)
// 	// }

// 	// go func() {
// 	// 	wg.Wait()
// 	// 	close(ch)
// 	// }()
// 	// for range urls {
// 	// 	fmt.Print(<-ch)
// 	// }

// 	ch := make(chan string)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Proccessdata("https://www.google.com", ch)
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Proccessdata("https://www.facebook.com", ch)
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Proccessdata("https://www.google.com", ch)
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Proccessdata("https://www.facebook.com", ch)
// 	}()

// 	go func() {
// 		wg.Wait()
// 	}()

// 	for i := 0; i < 4; i++ {
// 		fmt.Println(i, <-ch)
// 	}
// }

package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Connection closed")
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
