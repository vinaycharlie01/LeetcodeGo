package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var a []int
	var b []int
	temp := head
	for temp != nil {
		if temp.Val < x {
			a = append(a, temp.Val)
		} else {
			b = append(b, temp.Val)
		}
		temp = temp.Next
	}
	// sort.Ints(a)
	a = append(a, b...)
	temp2 := head
	i := 0
	for temp2 != nil {
		temp2.Val = a[i]
		i++
		temp2 = temp2.Next
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := new(ListNode)
	head := sentinel
	carry, val := 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		val = carry + getVal(l1) + getVal(l2)
		carry = val / 10
		val = val % 10

		head.Next = &ListNode{val, nil}
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

func deleteDuplicates(head *ListNode) *ListNode {
	temp := head
	if head == nil {
		return head
	}
	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			// new:=temp.Next.Next
			// temp.Next=nil
			// temp.Next=new
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head

	//  temp := llist
	// for temp.next != nil {
	//     if temp.data == temp.next.data {
	//         new := temp.next.next
	//         temp.next= nil
	//         temp.next = new
	//     }else{
	//         temp = temp.next
	//     }
	// }
	// return llist
}

func main() {
	a := []int{1, 1, 2, 3, 3}
	// [val,add]
	Link := new(ListNode)
	temp := Link
	for _, v := range a {
		// [0:50]-[1:100]--[1:200 ]--[val:]
		newNode := &ListNode{Val: v}
		temp.Next = newNode
		temp = temp.Next
	}
	fmt.Println(Display(deleteDuplicates(Link.Next)))
}

// func main() {

// a := []int{1, 4, 3, 0, 2, 5, 2}
// Link := new(ListNode)
// temp := Link
// for _, v := range a {
// 	temp.Next = &ListNode{v, nil}
// 	temp = temp.Next
// }
// k := 3
// fmt.Println(Display(partition(Link.Next, k)))
// }
