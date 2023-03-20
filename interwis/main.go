package main

import (
	"fmt"
	"strconv"
	"time"
)

func trafficSignalMessage(signalColor string) func(func() string) {
	// return func(message func() string) {
	// 	fmt.Println(signalColor, ":", message())
	// }
	return func(f func() string) {
		fmt.Println(signalColor, " :", f())
	}
}

type Decorator interface {
	Message() string
}
type name struct {
	name string
}

func (receiver name) Message() string {
	return receiver.name
}

func stopMessage() string {
	return "STOP"
}

func slowDownMessage() string {
	return "SLOW DOWN"
}

func goMessage() string {
	return "GO"
}

func Encode(a string) {
	var i int
	var count int
	b := ""
	for i < len(a)-1 {
		if a[i] == a[i+1] {
			count++
		}
		if a[i] != a[i+1] {
			if !(a[i] >= '1' && a[i] <= '9') {
				val := strconv.Itoa(count + 1)
				b = b + (string(a[i]) + val)
			}
			count = 0
		}
		i++
		if i == len(a)-1 {
			if a[i-1] == a[len(a)-1] && !(a[i-1] >= '1' && a[i-1] <= '9') {
				val := strconv.Itoa(count + 1)
				b = b + string(a[i-1]) + val
			}
		}
		if a[i] >= '1' && a[i] <= '9' {
			b = b + string(a[i])
		}
	}

	// ’a1b224c3d2b2c1a11’
	fmt.Println(b)
}

func main() {

	stopDecorator := trafficSignalMessage("RED")
	slowDownDecorator := trafficSignalMessage("YELLOW")
	goDecorator := trafficSignalMessage("GREEN")

	b := name{"STOP"}
	b1 := name{"SLOW DOWN"}
	var s Decorator = b
	var slow Decorator = b1
	stopDecorator(s.Message)
	slowDownDecorator(slow.Message)
	goDecorator(goMessage)

	interval := time.Second * 2
	timer := time.NewTimer(interval)
	for {
		select {
		case <-timer.C:
			fmt.Println("******")
			fmt.Println("*****")
			fmt.Println("****")
			fmt.Println("***")
			fmt.Println("**")
			fmt.Println("*")

		}
		timer.Reset(interval)
	}

}
