package main

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

// func main() {

// 	a := []int{1, 4, 3, 0, 2, 5, 2}
// 	Link := new(ListNode)
// 	temp := Link
// 	for _, v := range a {
// 		temp.Next = &ListNode{v, nil}
// 		temp = temp.Next
// 	}
// 	k := 3
// 	fmt.Println(Display(partition(Link.Next, k)))
// }
