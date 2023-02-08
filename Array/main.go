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

// func main() {
// 	Userinput := [][][]int{{{1, 2, 2, 1}, {2, 2}}, {{4, 9, 5}, {9, 4, 9, 8, 4}}, {{1, 2, 3, 2, 4, 5, 6, 7, 8, 9}, {1}}, {{1}, {0}}}
// 	Output := [][]int{{2}, {4, 9}, {1}, {}}
// 	var a [][]int
// 	for _, v := range Userinput {
// 		Result := intersection(v[0], v[1])
// 		a = append(a, Result)
// 	}
// 	fmt.Println(a)

// 	var map2 = map[int]int{}
// 	for _, v := range a {
// 		for _, val := range v {
// 			map2[val]++
// 		}
// 	}
// 	var map3 = map[int]int{}
// 	for _, v := range Output {
// 		for _, val := range v {
// 			map3[val]++
// 		}
// 	}

// 	fmt.Println(map2, map3)
// 	// a := map[int]int{1: 10, 0: 30, 3: 30}
// 	// b := map[int]int{0: 30, 3: 30, 1: 10}
// 	// fmt.Println(reflect.DeepEqual(a, b))
// }
