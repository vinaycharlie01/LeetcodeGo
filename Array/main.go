package main

func intersection(nums1 []int, nums2 []int) []int {
	Hashmap := make(map[int]int)
	for i, val := range nums2 {
		Hashmap[val] = i
	}
	hashmap2 := make(map[int]int)
	for _, val := range nums1 {
		_, ok := Hashmap[val]
		if ok {
			hashmap2[val]++
		}
	}
	var a []int
	for i, _ := range hashmap2 {
		a = append(a, i)
	}
	return a
}
