package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	head := sentinel
	carry, val := 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		val = carry + getVal(l1) + getVal(l2)
		carry = val / 10
		val = val % 10
		NewNode := &ListNode{val, nil}
		head.Next = NewNode
		head = head.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return sentinel.Next
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	hashmap := make(map[int]*ListNode)
	count := 0
	for temp != nil {
		count++
		hashmap[count] = temp
		temp = temp.Next
	}
	mid := count - n - 1
	if count == n {
		return head.Next
	}
	hashmap[mid].Next = hashmap[mid].Next.Next
	return head
}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	temp := head
	count := 1
	for temp.Next != nil {
		count++
		temp = temp.Next
	}
	mid := count - n
	temp2 := head
	for i := 0; i < mid; i++ {
		temp2 = temp2.Next
	}
	temp2.Next = temp2.Next.Next
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	L := l1
	R := l2
	NewList := new(ListNode)
	temp := NewList
	for L != nil && R != nil {
		if L.Val < R.Val {
			temp.Next = L
			L = L.Next
			temp = temp.Next
		} else {
			temp.Next = R
			R = R.Next
			temp = temp.Next
		}
	}
	if L != nil {
		temp.Next = L
	} else {
		temp.Next = R
	}
	return NewList.Next

}

func MergeSort(arr []*ListNode) *ListNode {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 0 {
		return nil
	}
	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	L := left
	R := right
	NewList := new(ListNode)
	temp := NewList
	for L != nil && R != nil {
		if L.Val < R.Val {
			temp.Next = L
			L = L.Next
			temp = temp.Next
		} else {
			temp.Next = R
			R = R.Next
			temp = temp.Next
		}
	}
	if L != nil {
		temp.Next = L
	} else {
		temp.Next = R
	}
	return NewList.Next

}

func mergeKLists(arr []*ListNode) *ListNode {
	// res := MergeSort(lists)
	// return res
	if len(arr) <= 1 {
		return arr[0]
	}
	mid := len(arr) / 2
	L := MergeSort(arr[:mid])
	R := MergeSort(arr[mid:])
	return mergeTwoLists(L, R)
}
func Display(head *ListNode) []int {
	var a []int
	temp := head
	for temp != nil {
		a = append(a, temp.Val)
		temp = temp.Next
	}
	return a
}

func swapPairs(head *ListNode) *ListNode {
	temp := head
	if head == nil {
		return head
	}
	var a []int
	for temp != nil {
		a = append(a, temp.Val)
		temp = temp.Next
	}
	for i := 0; i < len(a)-1; i += 2 {
		a[i], a[i+1] = a[i+1], a[i]
	}
	temp2 := head
	i := 0
	for temp2 != nil {
		temp2.Val = a[i]
		i++
		temp2 = temp2.Next
	}
	return head
}

func swapPairs2(head *ListNode) *ListNode {
	temp := head
	if head == nil {
		return head
	}
	i := 0
	for temp.Next != nil {
		if i%2 == 0 {
			temp.Val, temp.Next.Val = temp.Next.Val, temp.Val
		}
		i++
		temp = temp.Next
	}
	return head
}

// func rotateRight(head *ListNode, k int) *ListNode {

// 	for i := 0; i <= k; i++ {
// 		temp := head
// 		for temp != nil {
// 			temp.Val, temp.Next.Val = temp.Next.Val, temp.Val
// 			temp = temp.Next
// 		}
// 	}
// }

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	// Create a ring of integers
// 	for i := 0; i < k; i++ {
// 		temp := head
// 		for temp.Next != nil {
// 			temp.Val, temp.Next.Val = temp.Next.Val, temp.Val
// 			temp = temp.Next
// 		}
// 	}
// }

func rotateSlice(s []int, n int) []int {
	if n > 0 {
		n = n % len(s)
		return append(s[len(s)-n:], s[:len(s)-n]...)
	}
	n = -n % len(s)
	return append(s[n:], s[:n]...)
}

func Rotate(head *ListNode, tail *ListNode, s int, n int) *ListNode {
	if n > 0 {
		tail.Next = head
		n = n % s
		temp := head
		i := 0
		for i < n {
			i++
			temp = temp.Next
		}
		hea := temp.Next
		temp.Next = nil
		return hea
	}
	tail.Next = head
	n = n % s
	temp := head
	i := 0
	for i < n {
		i++
		temp = temp.Next
	}
	hea := temp.Next
	temp.Next = nil
	return hea
}

// func rotateRight(head *ListNode, k int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	// var s []int
// 	temp := head
// 	lenth := 0
// 	for temp != nil {
// 		lenth++
// 		// s = append(s, temp.Val)
// 		temp = temp.Next
// 	}
// 	val := rotateSlice(head, lenth, k)
// 	temp2 := head
// 	i := 0
// 	for temp2 != nil {
// 		temp2.Val = val.Val
// 		i++
// 		temp2 = temp2.Next
// 	}
// 	return head
// }

type SortedMap struct {
	data   []int
	adress *ListNode
}

func Recone(a string, dp map[string]int) map[string]int {
	if len(a) > 0 {
		dp[string(a[len(a)-1:])]++
		a = a[:len(a)-1]
		return Recone(a, dp)
	}
	return dp
}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	for head != nil {
// 		hashmap[head.Val]++
// 		head = head.Next
// 	}
// 	fmt.Println(hashmap)
// 	var a []int
// 	for i, val := range hashmap {
// 		if val == 1 {
// 			a = append(a, i)
// 		}
// 	}
// 	sort.Ints(a)
// 	i := 0
// 	newNode := new(ListNode)
// 	temp := newNode
// 	for i < len(a) {
// 		New := ListNode{a[i], nil}
// 		temp.Next = &New
// 		i++
// 		temp = temp.Next
// 	}
// 	return newNode.Next

// 	// return head

// }

func PrifixSum(a []int) {
	s := make([]int, len(a))
	for i := 1; i < len(a); i++ {
		s[i] = a[i-1] + s[i-1]
	}
	fmt.Println(s)
}

/*

Input: banned = [1,6,5], n = 5, maxSum = 6
Output: 2
Explanation: You can choose the integers 2 and 4.
2 and 4 are from the range [1, 5], both did not appear in banned, and their sum is 6, which did not exceed maxSum.
Example 2:

Input: banned = [1,2,3,4,5,6,7], n = 8, maxSum = 1
Output: 0
Explanation: You cannot choose any integer while following the mentioned conditions.
Example 3:

Input: banned = [11], n = 7, maxSum = 50
Output: 7
Explanation: You can choose the integers 1, 2, 3, 4, 5, 6, and 7.
They are from the range [1, 7], all did not appear in banned, and their sum is 28, which did not exceed maxSum.

*/

// func maxCount(banned []int, n int, maxSum int) {

// 	hashmap := make(map[int]int)
// 	for i, v := range banned {
// 		hashmap[v] = i
// 	}
// 	for i := 1; i < n; i++ {
// 		_, ok := hashmap[i]
// 		if !ok {
// 			fmt.Println(i)
// 			j++
// 		}
// 	}
// 	fmt.Println(j)
// 	//fmt.Println(count)
// 	// fmt.Println(hashmap)

// }

func main() {
	fmt.Println(Recone("Bharath", map[string]int{}))
	// a := []int{1, 6, 5}
	// maxCount(a, 5, 6)
	// PrifixSum(a)
	// s := map[int]string{1: "fnhfh", 2: "fnfn", 3: "fnfn", 4: "fjf", 5: "fdndj"}
	// delete(s, 1)
	// fmt.Println(s)

	//[4,5,1,2,3
	// fmt.Println("This is input ", a)
	// Link := new(ListNode)
	// temp := Link
	// for _, v := range a {
	// 	temp.Next = &ListNode{v, nil}
	// 	temp = temp.Next
	// }
	// fmt.Println(Display(Rotate(Link, len(a), 2)))

	// fmt.Println("OutPut", Display(swapPairs2(Link.Next)))
	// fmt.Println("result", Display(removeNthFromEnd3(Link.Next, 1)))
	// fmt.Println("want: ", "[3,7,9,5,8,2,3,0,0,0,3]")

	// a := []int{1, 2, 3}

	// for i := 0; i < len(a)-1; i += 2 {
	// 	// fmt.Println(i, i+1)
	// 	a[i], a[i+1] = a[i+1], a[i]
	// }
	// fmt.Println(a)
	// 	Input: head = [1,2,3,4,5], k = 2
	// Output: [4,5,1,2,3]
	// s := []int{0, 1, 2}
	// Rotate the slice to the right by 2
	// fmt.Println(rotateSlice(s, 4))

	// Print the rotated slice
	// fmt.Println(s)
}
