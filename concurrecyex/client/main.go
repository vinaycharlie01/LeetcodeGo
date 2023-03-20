// package main

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// )

// func MustCopy(in io.Writer, src io.Reader) {
// 	_, err := io.Copy(in, src)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	MustCopy(os.Stdout, conn)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Fun1(a string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; true; i++ {
// 		fmt.Println(a)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {

// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go Fun1("This is my first task", &wg)
// 	go Fun1("This is my second task", &wg)

// 	wg.Wait()

// }

package main

import (
	"fmt"
)

// CartItem interface
type CartItem interface {
	GetID() int
	GetTitle() string
	GetPrice() float64
	SetTitle(title string)
	SetPrice(price float64)
}

// Book struct
type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

// GetID method for Book
func (b *Book) GetID() int {
	return b.ID
}

// GetTitle method for Book
func (b *Book) GetTitle() string {
	return b.Title
}

// GetPrice method for Book
func (b *Book) GetPrice() float64 {
	return b.Price
}

// SetTitle method for Book
func (b *Book) SetTitle(title string) {
	b.Title = title
}

// SetPrice method for Book
func (b *Book) SetPrice(price float64) {
	b.Price = price
}

// Electronics struct
type Electronics struct {
	ID    int
	Title string
	Brand string
	Price float64
}

// GetID method for Electronics
func (e *Electronics) GetID() int {
	return e.ID
}

// GetTitle method for Electronics
func (e *Electronics) GetTitle() string {
	return e.Title
}

// GetPrice method for Electronics
func (e *Electronics) GetPrice() float64 {
	return e.Price
}

// SetTitle method for Electronics
func (e *Electronics) SetTitle(title string) {
	e.Title = title
}

// SetPrice method for Electronics
func (e *Electronics) SetPrice(price float64) {
	e.Price = price
}

// Cart struct
type Cart struct {
	items []CartItem
}

// Create method for Cart
func (c *Cart) Create(item CartItem) {
	c.items = append(c.items, item)
}

// Read method for Cart
func (c *Cart) Read(id int) CartItem {
	for _, item := range c.items {
		if item.GetID() == id {
			return item
		}
	}
	return nil
}

// Update method for Cart
func (c *Cart) Update(id int, title string, price float64) {
	for _, item := range c.items {
		if item.GetID() == id {
			item.SetTitle(title)
			item.SetPrice(price)
			break
		}
	}
}

// Delete method for Cart
func (c *Cart) Delete(id int) {
	for i, item := range c.items {
		if item.GetID() == id {
			c.items = append(c.items[:i], c.items[i+1:]...)
			break
		}
	}
}

func main() {
	// Create a new cart
	cart := &Cart{}

	// Add some items to the cart
	book := &Book{1, "The Great Gatsby", "F. Scott Fitzgerald", 9.99}
	cart.Create(book)

	electronics := &Electronics{2, "Echo Dot", "Amazon", 49.99}
	cart.Create(electronics)

	// Update the price of the book
	cart.Update(1, "The Great Gatsby (Paperback)", 12.99)

	// Remove the Echo Dot from the cart
	cart.Delete(2)

	// Print the remaining items in the cart
	for _, item := range cart.items {
		fmt.Printf("%s: $%.2f\n", item.GetTitle(), item.GetPrice())
	}
}
