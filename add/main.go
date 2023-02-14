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
