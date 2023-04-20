package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

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

func containsDuplicate(nums []int) bool {
	hashmap := make(map[int]int)
	for _, v := range nums {
		_, ok := hashmap[v]
		if ok {
			return true
		}
		hashmap[v]++
	}
	return false
}

func findKthLargest(nums []int, k int) int {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sort.Ints(nums)
	}()
	wg.Wait()
	return nums[len(nums)-k]
}

func wiggleSort(nums []int) {
	// mid := len(nums) / 2
	// l1 := nums[:mid]
	// sort.Sort(sort.Reverse(sort.IntSlice(l1)))
	// r1 := nums[mid:]
	// sort.Sort(sort.Reverse(sort.IntSlice(r1)))
	// fmt.Println(l1, r1)
	// L := 0
	// R := 0
	// var a []int
	// for L < len(l1) && R < len(r1) {
	// 	a = append(a, r1[R])
	// 	R++
	// 	a = append(a, l1[L])
	// 	L++
	// }
	// copy(nums, a)
	// fmt.Println(a)
}

func wiggleSort2(nums []int) {
	mid := len(nums) / 2
	l1 := nums[:mid]
	sort.Sort(sort.Reverse(sort.IntSlice(l1)))
	r1 := nums[mid:]
	sort.Sort(sort.Reverse(sort.IntSlice(r1)))
	fmt.Println(l1, r1)
	L := 0
	R := 0
	var a []int
	for L < len(l1) && R < len(r1) {
		a = append(a, r1[R])
		R++
		a = append(a, l1[L])
		L++
	}
	copy(nums, a)
	fmt.Println(a)
}

func findTheArrayConcVal(nums []int) int64 {
	L := 0
	R := len(nums) - 1
	var a int
	for L <= R {
		if L == R {
			a += nums[L]
		} else {
			val := strconv.Itoa(nums[L]) + strconv.Itoa(nums[R])
			s, _ := strconv.Atoi(val)
			a += s
		}
		L++
		R--

	}
	return int64(a)
}

func findMaxConsecutiveOnes(nums []int) {
	// hashmap := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		}
	}
	fmt.Println(count)
}

/*Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).*/
func removeElement(nums []int, val int) int {
	i := 0
	for true {
		if i == len(nums) {
			break
		}
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
		i += 1
	}
	return len(nums)

}

func intersect(nums1 []int, nums2 []int) []int {
	hashmap := make(map[int]int)
	for _, v := range nums1 {
		hashmap[v]++
	}
	var a []int
	for _, v := range nums2 {
		val, ok := hashmap[v]
		if ok {
			a = append(a, val)
		}
	}
	return a
}

func conarr(nums1 [][]int, nums2 [][]int, ch chan [][]int) {
	defer close(ch)
	var a [][]int
	L := 0
	R := 0
	for L < len(nums1) && R < len(nums2) {
		if nums1[L][0] == nums2[R][0] {
			a = append(a, []int{nums1[L][0], nums1[L][1] + nums2[L][1]})
		} else if nums1[L][0] < nums2[R][0] {
			a = append(a, nums1[L], nums2[R])
		} else {
			a = append(a, nums2[R], nums1[L])
		}
		L++
		R++
	}
	if L < len(nums1) {
		a = append(a, nums1[L:]...)
	} else {
		a = append(a, nums2[R:]...)
	}
	ch <- a
}

func chanarr(a [][]int, ch chan [][]int) {
	defer close(ch)
	hm := make(map[int]int)
	for _, v := range a {
		hm[v[0]] += v[1]
	}
	var c [][]int
	for i, v := range hm {
		c = append(c, []int{i, v})
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i][0] < c[j][0]
	})
	ch <- c
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	ch := make(chan [][]int)
	var wg sync.WaitGroup
	// go onarr(nums1, nums2, ch)
	wg.Add(2)
	go func() {
		defer wg.Done()
		conarr(nums1, nums2, ch)
	}()
	a := <-ch
	ch2 := make(chan [][]int)
	go func() {
		defer wg.Done()
		chanarr(a, ch2)
	}()
	b := <-ch2
	// fmt.Println(b)

	wg.Wait()
	// var a [][]int
	// L := 0
	// R := 0
	// for L < len(nums1) && R < len(nums2) {
	// 	if nums1[L][0] == nums2[R][0] {
	// 		a = append(a, []int{nums1[L][0], nums1[L][1] + nums2[L][1]})
	// 	} else if nums1[L][0] < nums2[R][0] {
	// 		a = append(a, nums1[L], nums2[R])
	// 	} else {
	// 		a = append(a, nums2[R], nums1[L])
	// 	}
	// 	L++
	// 	R++
	// }
	// if L < len(nums1) {
	// 	a = append(a, nums1[L:]...)
	// } else {
	// 	a = append(a, nums2[R:]...)
	// }
	// a := <-ch
	// hm := make(map[int]int)
	// for _, v := range a {
	// 	hm[v[0]] += v[1]
	// }
	//var c [][]int
	// for i, v := range hm {
	// 	c = append(c, []int{i, v})
	// }
	// sort.Slice(c, func(i, j int) bool {
	// 	return c[i][0] < c[j][0]
	// })

	// fmt.Println(c)
	return b
}

// Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
// Output: 6
// Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).
func countFairPairs(nums []int, lower int, upper int) {
	// sort.IntSlice(nums).Sort()
	// // fmt.Println(nums)
	// L := 0
	// R := len(nums) - 1
	// for L <= R {
	// 	mid := (L + R) / 2

	// 	if nums[mid]

	// 	L++
	// 	R--
	// }
	// fmt.Println(hashmap)

}

func productExceptSelf(nums []int) {
	// prifix := make([]int, len(nums)+1)
	// fmt.Println(prifix)
	// fmt.Println(nums)
	// for i := 0; i <= len(nums); i++ {
	// 	prifix[i] = prifix[i] + nums[i-1]
	// }
	// fmt.Println(prifix)

	prefix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prefix[i] = nums[i]
		if i > 0 {
			prefix[i] = prefix[i] + prefix[i-1]
		}
	}
	fmt.Println(prefix)
}

// var m = make(map[any]any)

// func foo1[T any](x T) {
// 	type K struct{}
// 	m[K{}] = x
// }

func Max(arr []int) int {
	sort.IntSlice(arr).Sort()
	return arr[len(arr)-1]
}
func Min(arr []int) int {
	sort.IntSlice(arr).Sort()
	return arr[0]
}

/*
Input: nums = [10,4,8,3]
Output: [15,1,11,22]
Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].
*/

func leftRigthDifference(nums []int) []int {
	fmt.Println(nums)
	prifix := make([]int, len(nums)+1)
	for i := 1; i < len(nums); i++ {
		prifix[i] = prifix[i-1] + nums[i-1]
	}
	fmt.Println(prifix)
	b := nums
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	prifix2 := make([]int, len(b)+1)
	for i := 1; i < len(b); i++ {
		prifix2[i] = prifix2[i-1] + b[i-1]
	}
	prifix = prifix[:len(prifix)-1]
	prifix2 = prifix2[:len(prifix2)-1]
	fmt.Println(prifix, prifix2)

	start := 0
	end := len(nums) - 1
	var a []int
	for start < len(nums) && end >= 0 {
		a = append(a, int(math.Abs(float64(prifix[start]-prifix2[end]))))
		start++
		end--
	}
	return a
}

func Add(a string) {

	var right int
	var left int
	b := ""
	// "abb24ccc8ddbbca1"
	for right < len(a)-1 {
		if a[right] == a[right+1] {
			left++
			right++
		}
		if a[right] != a[right+1] {
			if !(a[right] >= '1' && a[right] <= '9') {
				val := strconv.Itoa(left + 1)
				b = b + (string(a[right]) + val)
			} else {
				b = b + string(a[right])
			}
			left = 0
			right++
		}
		if right == len(a)-1 {
			b = b + string(a[len(a)-1])
		}
	}
	// ’a1b224c3d2b2c1a11’
	fmt.Println(b)
}

//   s := "abb24ccc8ddbbca1"
// fmt.Println(encode(s)) // output: "a1b224c3d2b2c1a11"
//
func Encode(a string) {
	var right int
	var left int
	b := ""
	for right < len(a)-1 {
		if a[right] == a[right+1] {
			left++
		}
		if a[right] != a[right+1] {
			if !(a[right] >= '1' && a[right] <= '9') {
				val := strconv.Itoa(left + 1)
				b = b + (string(a[right]) + val)
			}
			left = 0
		}
		right++
		if right == len(a)-1 {
			if a[right-1] == a[len(a)-1] && !(a[right-1] >= '1' && a[right-1] <= '9') {
				val := strconv.Itoa(left + 1)
				b = b + string(a[right-1]) + val
			}
		}
		if a[right] >= '1' && a[right] <= '9' {
			b = b + string(a[right])
		}
	}

	// ’a1b224c3d2b2c1a11’
	fmt.Println(b)
}

func passThePillow(n int, time int) {
	i := 0
	m := make([]int, n+1)
	count := 0

	for i < time+1 {
		if count > len(m) {
			fmt.Println(count)
			count--
		}
		count++
		i++
	}
	fmt.Println(m)
}

func rotate(s []int, n int) {
	// [1, 2, 3, 4, 5, 6, 7] 7.3
	var b []int

	if n > 0 {
		n = n % len(s)
		fmt.Println(n)
		b = append(s[len(s)-n:], s[:len(s)-n]...)
	} else {
		n = -n % len(s)
		b = append(s[n:], s[:n]...)
	}
	fmt.Println(b)
}
func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, -1)
	// passThePillow(6, 10)
	// a := []int{10, 4, 8, 3}
	// fmt.Println(Max(a))
	// foo1(123)
	// foo1(true)
	// fmt.Println(len(m))

	// a := [][]int{{1, 4}, {1, 6}, {5, 5}}
	// b := [][]int{{1, 3}, {4, 3}}
	// fmt.Println(mergeArrays(a, b))
	// a := []int{0, 1, 7, 4, 4, 5}
	// a := []int{1, 2, 3, 4}
	// productExceptSelf(a)
	// countFairPairs(a, 3, 6)

	// nums1 = [1,2,2,1], nums2 = [2,2]
	// a := []int{4, 9, 5}
	// b := []int{9, 4, 9, 8, 4}
	// fmt.Println(intersect(a, b))
	// findDisappearedNumbers(a)
}

// func main() {

// 	a := []int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3}
// 	// b := []int{1, 3}
// 	// wiggleSort2(a)
// 	// wiggleSort(b)

// }

func groupAnagrams(a []string) [][]string {
	dp := make(map[string][]string)
	for i1 := 0; i1 < len(a); i1++ {
		b := []byte(a[i1])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		// fmt.Println([]byte(a[i1]))
		dp[string(b)] = append(dp[string(b)], a[i1])
		// fmt.Println(b[i1])
	}

	var s [][]string
	for _, v := range dp {
		s = append(s, v)
	}
	return s
}

/*
Input: ops = ["5","2","C","D","+"]
Output: 30
Explanation:
"5" - Add 5 to the record, record is now [5].
"2" - Add 2 to the record, record is now [5, 2].
"C" - Invalidate and remove the previous score, record is now [5].
"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
The total sum is 5 + 10 + 15 = 30.

*/
func CalPoints(nums []string) int {
	var res2 []int
	for i := 0; i < len(nums); i++ {
		res, err := strconv.Atoi(nums[i])
		if err != nil {
			if nums[i] == "D" {
				r := res2[len(res2)-1]
				res2 = append(res2, 2*r)
			}
			if nums[i] == "C" && i > 0 {
				res2 = res2[:len(res2)-1]
				// res2 = append(res2, )
			}
			if nums[i] == "+" && i > 1 {
				fmt.Println(res2[len(res2)-1], res2[len(res2)-2])
				res2 = append(res2, res2[len(res2)-1]+res2[len(res2)-2])
			}
		} else {
			res2 = append(res2, res)
		}
	}
	fmt.Println(res2)
	var res int
	for _, v := range res2 {
		res += v
	}
	fmt.Println(res)
	return res
}

func IsMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var a bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			a = true
		} else {
			a = false
			break
		}
	}
	if a != true {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] >= nums[i+1] {
				a = true
			} else {
				a = false
				break
			}
		}
	}
	fmt.Println(a)
	return a
}
