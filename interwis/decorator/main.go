package main

import "fmt"

func TrafficSignalMessage(f func(string)) func(string) {
	return func(color string) {
		fmt.Printf("%s : ", string(color))
		f(color)
	}
}

func StopMessage(color string) {
	fmt.Println("STOP")
}

func SlowDownMessage(color string) {
	fmt.Println("SLOW DOWN")
}

func GoMessage(color string) {
	fmt.Println("GO")
}

func main() {
	stop := TrafficSignalMessage(StopMessage)
	slowDown := TrafficSignalMessage(SlowDownMessage)
	goMsg := TrafficSignalMessage(GoMessage)

	stop("red")
	slowDown("yellow")
	goMsg("green")
}
