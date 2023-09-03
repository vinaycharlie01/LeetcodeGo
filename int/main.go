package main

import (
	"fmt"
	"strconv"

	"github.com/emirpasic/gods/sets/hashset"
	"golang.org/x/exp/maps"
)

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

func MinimumSum(n int, k int) {

	hashset := map[int]struct{}{}
	for i := 1; len(hashset) != n; i++ {
		_, ok := hashset[k-i]
		if !ok {
			hashset[i] = struct{}{}
		}

	}
	r := maps.Keys(hashset)
	fmt.Println(r)
	var a int
	for _, v := range r {
		a += v
	}
	fmt.Println(a)

}

func MinimumSum2(n int, k int) {
	hashset := hashset.New()
	for i := 1; hashset.Size() != n; i++ {
		fmt.Println(i)
		if !hashset.Contains(k - i) {
			hashset.Add(i)
		}
	}
	var a int
	for _, v := range hashset.Values() {
		a += v.(int)
	}
	// fmt.Println(a)

}

func CountSymmetricIntegers(low int, high int) int {
	var count int
	for i := low; i <= high; i++ {
		b := strconv.Itoa(i)
		n := len(b)
		l := n / 2
		// fmt.Println(b)
		// fmt.Println(l)
		if n%2 == 1 {
			continue
		}
		sum := 0
		for i := 0; i < l; i++ {
			num, _ := strconv.Atoi(string((b[i])))
			sum += num
		}
		sum2 := 0
		for i := l; i < n; i++ {
			num, _ := strconv.Atoi(string((b[i])))
			sum2 += num
		}
		if sum == sum2 {
			count++
		}
	}
	return count
}

/*
Given two positive integers low and high, this function counts the number of symmetric integers in the range [low, high].
An integer x consisting of 2 * n digits is symmetric if the sum of the first n digits of x is equal to the sum of the last n digits of x. Numbers with an odd number of digits are never symmetric.

@param low: The lower bound of the range (inclusive).
@param high: The upper bound of the range (inclusive).

@return: The count of symmetric integers in the range [low, high].
*/
func CountSymmetricIntegers2(low int, high int) int {
	var count int
	for i := low; i <= high; i++ {
		lenth := 0
		temp := i
		for temp > 0 {
			temp /= 10
			lenth++

		}
		if lenth%2 == 1 {
			continue
		}
		temp = i
		sum1 := 0
		sum2 := 0
		for j := 0; j < lenth; j++ {
			digit := temp % 10
			temp /= 10
			if j < lenth/2 {
				sum1 += digit
			} else {
				sum2 += digit
			}
		}
		if sum1 == sum2 {
			count++
		}
	}
	return count
}

func main() {

	// countSymmetricIntegers(100, 1782)

	// set := hashset.New()   // empty
	// set.Add(1)             // 1
	// set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	// set.Remove(4)          // 5, 3, 2, 1 (random order)
	// set.Remove(2, 3)       // 1, 5 (random order)
	// set.Contains(1)        // true
	// set.Contains(1, 5)     // true
	// set.Contains(1, 6)     // false
	// _ = set.Values()       // []int{5,1} (random order)
	// set.Clear()            // empty
	// set.Empty()            // true
	// set.Size()

	// a := []int{1, 2, 4, 5, 7, 9, 12, 14, 16}
	// MinimumSum2(5, 4)

}
