package main

import "fmt"

func Hash(arr []int, target int) {
	hashmap := make(map[int]int)
	var a []int
	for i, v := range arr {
		val, ok := hashmap[target-v]
		if ok {
			a = append(a, val, i)
		}
		hashmap[v] = i

	}
	fmt.Println(a)
}

func main() {

}
