package main

func Subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, num := range nums {
		for _, subset := range res {

			newSubset := append([]int{}, subset...)
			// fmt.Println(newSubset)
			newSubset = append(newSubset, num)
			// fmt.Println(num)
			// Append the new subset to the result list
			res = append(res, newSubset)
			// fmt.Println(newSubset)
		}
	}
	// fmt.Println(res)
	// Return the result list
	return res
}
