package main

import "golang.org/x/exp/slices"

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

// genera
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

/*
n = len(nums)

	distinct_elements = len(set(nums))
	count = 0
	left = 0
	right = 0
	counter = Counter()

	while right < n:
	    counter[nums[right]] += 1
	    while len(counter) == distinct_elements:
	        counter[nums[left]] -= 1
	        if counter[nums[left]] == 0:
	            del counter[nums[left]]
	        left += 1
	        count += n - right
	    right += 1

	return count
*/
func CountCompleteSubarrays(nums []int) int {
	n := len(nums)
	distinct_elements := map[int]int{}
	count := 0
	counter := map[int]int{}
	for _, v := range nums {
		distinct_elements[v]++
	}
	left := 0
	right := 0
	for right < n {
		counter[nums[right]]++
		for len(counter) == len(distinct_elements) {
			counter[nums[left]]--
			if counter[nums[left]] == 0 {
				delete(counter, nums[left])
			}
			left++
			count += n - right
		}
		right += 1

	}
	return count
}

func main() {
	a := []int{5, 5, 5, 5}
	// fmt.Println(a)
	CountCompleteSubarrays(a)

}

func GenerateSubarrays2(nums []int) [][]int {
	result := [][]int{}
	Backtrack(nums, 0, []int{}, &result)
	return result
}

func Backtrack(nums []int, start int, current []int, result *[][]int) {
	// Add a copy of the current subarray to the result
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])    // Include the current element
		Backtrack(nums, i+1, current, result) // Recurse with next index
		current = current[:len(current)-1]    // Exclude the current element
	}
}

func alternatingSubarray(nums []int) int {
	n := len(nums)
	res, dp := -1, -1
	for i := 1; i < n; i+slices
		if dp > 0 && nums[i] == nums[i-2] {
			dp += 1
		} else {
			if nums[i] == nums[i-1]+1 {
				dp = 2
			} else {
				dp = -1
			}
		}
		if dp > res {
			res = dp
		}
	}
	return res
}
