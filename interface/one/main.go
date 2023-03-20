// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Fun1(a string) {
// 	for i := 0; i < 4; i++ {
// 		fmt.Println(i, a)
// 		time.Sleep(time.Second * 3)
// 	}
// }

// func main() {

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Fun1("this is my tast 1")
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Fun1("this is my tast 2")
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		Fun1("this is my tast 3")
// 	}()
// 	wg.Wait()

// }

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Hello World</H1>")
	fmt.Fprintln(w, "<H2>This is Vinay</H2>")
}

func main() {
    
	http.HandleFunc("/", handler)
	
	http.ListenAndServe(":8080", nil)
}