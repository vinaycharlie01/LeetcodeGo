package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Intersection(nums1 []int, nums2 []int) []int {
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

func ContainsDuplicate(nums []int) bool {
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

func FindKthLargest(nums []int, k int) int {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sort.Ints(nums)
	}()
	wg.Wait()
	return nums[len(nums)-k]
}

func WiggleSort(nums []int) []int {
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
	return a
}

func FindTheArrayConcVal(nums []int) int64 {
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

func FindMaxConsecutiveOnes(nums []int) (count int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		}
	}
	return count
}

/*
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/
func RemoveElement(nums []int, val int) int {
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

func Intersect(nums1 []int, nums2 []int) []int {
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

func MergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
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

func ProductExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prefix[i] = nums[i]
		if i > 0 {
			prefix[i] = prefix[i] + prefix[i-1]
		}
	}
	return prefix
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
func LeftRigthDifference(nums []int) []int {
	// fmt.Println(nums)
	prifix := make([]int, len(nums)+1)
	for i := 1; i < len(nums); i++ {
		prifix[i] = prifix[i-1] + nums[i-1]
	}
	// fmt.Println(prifix)
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
	// fmt.Println(prifix, prifix2)
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

func Rotate(s []int, n int) []int {
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
	return b
}

// func main() {

// 	a := []int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3}
// 	// b := []int{1, 3}
// 	// wiggleSort2(a)
// 	// wiggleSort(b)

// }

func GroupAnagrams(a []string) [][]string {
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

func MoveZeroes(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}

	}
	return nums
}

func Rotate2(matrix [][]int) [][]int {
	var b [][]int
	for i := 0; i < len(matrix); i++ {
		var a []int
		for j := len(matrix) - 1; j >= 0; j-- {
			a = append(a, matrix[j][i])
		}
		b = append(b, a)
	}
	copy(matrix, b)
	return b
}

func IsWinner(player1 []int, player2 []int) int {
	var count1 int
	var count2 int
	for i := 0; i < len(player1); i++ {
		j := i - 1
		j2 := i - 2
		if j >= 0 && player1[j] == 10 {
			count1 += 2 * player1[i]
		} else if j2 >= 0 && player1[j2] == 10 {
			count1 += 2 * player1[i]
		} else {
			count1 += player1[i]
		}
		if j >= 0 && player2[j] == 10 {
			count2 += 2 * player2[i]

		} else if j2 >= 0 && player2[j2] == 10 {
			count2 += 2 * player2[i]
		} else {
			count2 += player2[i]
		}

	}
	if count1 == count2 {
		return 0
	} else if count1 > count2 {
		return 1
	} else {
		return 2
	}
}

func maxProduct(words []string) int {
	var a int
	for i := 0; i < len(words); i++ {
		for j := len(words) - 1; j >= i; j-- {
			if !strings.ContainsAny(words[i], words[j]) {
				s1 := len(words[i]) * len(words[j])
				if s1 > a {
					a = s1
				}
			}
		}
	}
	return a
	// fmt.Println(a)
	// // sort.Ints(a1)
	// // if sort.Ints(a1); a1 != nil {
	// // 	fmt.Println(a1[len(a1)-1])
	// // } else {
	// // 	fmt.Println(0)
	// // }
}

func BinarySearch(a []int, target int) bool {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return true
		} else if target > a[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false

}

/*
There are n employees in a company, numbered from 0 to n - 1. Each employee i has worked for hours[i] hours in the company.
The company requires each employee to work for at least target hours.
You are given a 0-indexed array of non-negative integers hours of length n and a non-negative integer target.
Return the integer denoting the number of employees who worked at least target hours
*/
func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var count int
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			count++
		}
	}
	return count

}

func MakeRandomInts(n int) []int {
	rand.Seed(10)
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(n)
	}
	return ints
}

func RandomInts(n int) int {
	return rand.Intn(n)
}

func MinAbsoluteDifference(nums []int, x int) int {
	var a int = math.MaxInt
	fmt.Println(a)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(i-j))) >= x {
				b := int(math.Abs(float64(nums[i] - nums[j])))
				if b < a {
					a = b
				}
			}

		}
	}
	return a

}

func FindDuplicates(nums []int) []int {
	var duplicates []int
	for _, num := range nums {
		if nums[abs(num)-1] < 0 {
			duplicates = append(duplicates, num)
		} else {
			nums[abs(num)-1] = -nums[abs(num)-1]
		}
	}

	return duplicates
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/*
You are given a 0-indexed integer array nums. You have to find the maximum sum of a pair of numbers from nums such that the maximum digit in both numbers are equal.
Return the maximum sum or -1 if no such pair exists.

Example 1:

Input: nums = [51,71,17,24,42]
Output: 88
Explanation:
For i = 1 and j = 2, nums[i] and nums[j] have equal maximum digits with a pair sum of 71 + 17 = 88.
For i = 3 and j = 4, nums[i] and nums[j] have equal maximum digits with a pair sum of 24 + 42 = 66.
It can be shown that there are no other pairs with equal maximum digits, so the answer is 88.

Example 2:

Input: nums = [1,2,3,4]
Output: -1
Explanation: No pair exists in nums with equal maximum digits.
*/
func MaxSum(nums []int) {
	hashmap := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		var arr []int
		n := nums[i]
		var a int
		for n > 0 {
			r := n % 10
			arr = append(arr, r)
			a = a*10 + r
			n /= 10
			// arr = append(arr, r)
		}
		fmt.Println(a)
		// sort.Ints(arr)
		// fmt.Println(v)
		hashmap[nums[i]] = arr

		// fmt.Println(nums[i], a)
		// if reflect.DeepEqual(nums[i], a) {
		// 	fmt.Println(nums[i], a)
		// }

		// v := maps.Values(hashmap)
		//

		// hashmap[i] = arr

	}

	// v := maps.Values(hashmap)
	// fmt.Println(v)
	// // fmt.Println(slices.Contains(v, arr))
	// fmt.Println(slices.ContainsFunc(v, func(i []int) bool {
	// 	return reflect.DeepEqual(arr, i)
	// }))
}


func countSeniors(details []string) int {
	var count int
	for _, v := range details {
		val, _ := strconv.Atoi(v[11:13])
		if val > 60 {
			count++
		}
	}
	return count
}

func main() {
	MaxSum([]int{51, 71, 17, 24, 42, 17})

	// a := []int{1, 1, 2}
	// fmt.Println(FindDuplicates(a))
	// a := map[int]int{0: 0, 1: 1, 2: 2}

	// m := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// 	"four":  4,
	// }
	// n := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// 	"four":  6,
	// }

	// fmt.Println(time.Now().UTC().UTC())
	// maps.DeleteFunc(m, func(k string, v int) bool {
	// 	return v == 2
	// })
	// time.UTC()
	// fmt.Println(maps.Equal(m, n))
	// fmt.Println(maps.Keys(m))
	// fmt.Println(maps.Values(m))
	// // maps.Clear[]()
	// fmt.Println(reflect.DeepEqual(m, n))
	// fmt.Println(m)

	// fmt.Println(MinAbsoluteDifference([]int{330702844, 313481959, 239224564, 609763700, 170531905}, 0))
	// fmt.Println(rand.Intn(10))
	// fmt.Println(makeRandomInts(10))

	// fmt.Println(NumberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))

	// fmt.Println(strings.ContainsAny("vinay", "kummr"))

	// a := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	// maxProduct(a)
	// fmt.Println(strings.ContainsAny(a[0], a[4]))
	// a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// rotate2(a)

	// rotate([]int{1, 2, 3, 4, 5, 6, 7}, -1)
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
