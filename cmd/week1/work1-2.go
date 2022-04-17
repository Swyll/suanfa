package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//合并有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p1, p2 := list1, list2
	var p3, head *ListNode

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			if p3 == nil {
				p3, p1 = p1, p1.Next
				head = p3
			} else {
				p3, p3.Next, p1 = p1, p1, p1.Next
			}

			if p1 == nil {
				p3.Next = p2
				break
			}
		} else {
			if p3 == nil {
				p3, p2 = p2, p2.Next
				head = p3
			} else {
				p3, p3.Next, p2 = p2, p2, p2.Next
			}

			if p2 == nil {
				p3.Next = p1
				break
			}
		}
	}

	return head
}
