// package main

// import "fmt"

// func smallestEvenMultiple(n int) int {
// 	if n%2 == 1 {
// 		return n * 2
// 	} else {
// 		return n
// 	}

// }

// func main() {
// 	var n int
// 	fmt.Println("enter the number: ")
// 	fmt.Scanln(&n)
// 	fmt.Print(smallestEvenMultiple(n))
// }

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func Pattern(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if i <= j {
				fmt.Print("*", " ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Pyramid(r int) {
	// spc := r - 1
	// str := 1
	for i := 0; i < r; i++ {
		for j := 0; j < r-1-i; j++ {
			fmt.Print(" ", " ")
		}
		for k := 0; k < i+1+i; k++ {
			fmt.Print("*", " ")
		}
		fmt.Println()
	}

}

func subsetsWithDup(nums []int) {
	a := make(map[int][]int)
	s := make(map[int][]int)
	if len(nums) > 0 {
		for i := 0; i <= len(nums); i++ {
			b := []int{}
			b = append(b, nums[:i]...)
			a[i] = b
		}

		for i := len(nums) - 1; i > 0; i-- {
			s1 := []int{}
			s1 = append(s1, nums[i:]...)
			s[i] = s1
		}
		nums = nums[len(nums)-1:]
		subsetsWithDup(nums)
	}
	var a1 [][]int
	var a3 [][]int
	for _, v := range a {
		if len(v) > 0 {
			a1 = append(a1, v)
		} else {
			a3 = append(a3, v)
		}
	}
	for _, v := range s {
		if len(v) > 0 {
			a1 = append(a1, v)
		} else {
			a3 = append(a3, v)
		}

	}
	fmt.Println(a1)
	sort.Slice(a1, func(i, j int) bool {
		return a1[i][0] < a1[j][0]
	})
	for _, v := range a1 {
		a3 = append(a3, v)
	}
	// a3 = append(a3, a1...)
	fmt.Println(a3)
}

type Cart interface {
	add(a string) []string
}
type books struct {
	name []string
}

type Mobiles struct {
	name []string
}

func (b *books) add(a string) []string {
	b.name = append(b.name, a)
	return b.name
}
func (b *Mobiles) add(a string) []string {
	b.name = append(b.name, a)
	return b.name
}

type cloths struct {
	name []string
}

func (b *cloths) add(a string) []string {
	b.name = append(b.name, a)
	return b.name
}

func All(c Cart, s string) []string {
	return c.add(s)
}

// type SortBy []Type

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	a := books{
		name: []string{"GOBook", "Pythonbook", "Java"},
	}
	b := Mobiles{
		name: []string{"Realme", "iphome", "Mi"},
	}
	c := cloths{
		name: []string{"red", "blue", "yello"},
	}
	var s []Cart
	s = append(s, &a)
	s = append(s, &b)
	s = append(s, &c)

	fmt.Println(reflect.TypeOf(s[0]))
	fmt.Println(reflect.TypeOf(s[1]))
	var q any = []int{10, 20, 30}

	switch q.(type) {
	case string:
		fmt.Println("Stinrg")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("")
	}

	// a := []int{1, 2, 2, 3, 4, 5, 6}

	// a := []string{`nfjfj`, `jfnfhfj`, `fnfnfjhf`, `fnfhfj`, `fnfnfjhfj`}
	// fmt.Println(a)

	// a := [][]int{
	// 	{3, 1, 4},
	// 	{1, 5, 9},
	// 	{2, 6, 5},
	// 	{},
	// }

	// // sort the array by the first element of each subarray
	// sort.Slice(a, func(i, j int) bool {
	// 	if len(a[i]) > 1 {
	// 		return a[i][0] < a[j][0]
	// 	} else {
	// 		return false
	// 	}
	// })

	// // print the sorted array
	// fmt.Println(a)
	// subsetsWithDup(a)
	// subsetsWithDup(a)
	// Pyramid(10)
	// the empty struct struct{} is a special type that has zero fields, and thus occupies no memory

	// a := []interface{}{10, 20, "ujfufn"}
	// var s []int
	// var s2 []string
	// for _, v := range a {
	// 	val, ok := v.(int)
	// 	if ok {
	// 		s = append(s, val)
	// 	} else if val, ok := v.(string); ok {
	// 		s2 = append(s2, val)
	// 	}
	// }
	// fmt.Println(s, s2)
	// fmt.Println(str)
	// Pattern(10)
}

func ADD(a, b int) int {
	return a + b
}

func TwoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
		if len(result) > 0 {
			break
		}
	}
	return result
}
