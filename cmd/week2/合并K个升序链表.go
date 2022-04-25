package main

import "fmt"

/**
23. 合并K个升序链表

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
**/
func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3

	n5 := &ListNode{Val: 1}
	n6 := &ListNode{Val: 3}
	n7 := &ListNode{Val: 4}

	n5.Next = n6
	n6.Next = n7

	n8 := &ListNode{Val: 2}
	n9 := &ListNode{Val: 6}

	n8.Next = n9

	n := mergeKLists([]*ListNode{n1, n5, n8})

	p := n
	for {
		if p == nil {
			break
		}
		fmt.Println(p.Val)
		p = p.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	//给过来的数组可能有空数组，要考虑到该边界条件
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	length := len(lists)

	r1 := mergeKLists(lists[:length/2])
	r2 := mergeKLists(lists[length/2 : length])

	return merge(r1, r2)
}

func merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p1, p2 := list1, list2
	head := &ListNode{}
	tail := head

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail.Next, p1, tail = p1, p1.Next, p1

			if p1 == nil {
				tail.Next = p2
				break
			}
		} else {
			tail.Next, p2, tail = p2, p2.Next, p2

			if p2 == nil {
				tail.Next = p1
				break
			}
		}
	}

	return head.Next
}
