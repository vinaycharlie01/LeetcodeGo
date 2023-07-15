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

func generateSubarrays(nums []int) [][]int {
	subarrays := [][]int{}
	n := len(nums)

	for length := 0; length <= n; length++ {
		for start := 0; start+length <= n; start++ {
			subarray := make([]int, length)
			copy(subarray, nums[start:start+length])
			subarrays = append(subarrays, subarray)
		}
	}

	return subarrays
}
func generateSubarrays(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, start int, current []int, result *[][]int) {
	// Add a copy of the current subarray to the result
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])    // Include the current element
		backtrack(nums, i+1, current, result) // Recurse with next index
		current = current[:len(current)-1]    // Exclude the current element
	}
}
