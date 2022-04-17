package main

/**
设计实现双端队列。

实现 MyCircularDeque 类:

MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
**/

func main() {
}

type Node struct {
	val  int
	Next *Node
	Pre  *Node
}

type MyCircularDeque struct {
	head *Node
	tail *Node
	//head2    *Node
	//tail2    *Node
	length   int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head: nil,
		tail: nil,
		//head2:    nil,
		//tail2:    nil,
		length:   0,
		capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.capacity {
		return false
	}

	if this.head == nil {
		node := &Node{
			val: value,
		}
		this.head = node
		this.tail = node
	} else {
		node := &Node{
			val: value,
		}
		this.head.Pre, node.Next, this.head = node, this.head, node
	}

	this.length++

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.capacity {
		return false
	}

	if this.tail == nil {
		node := &Node{
			val: value,
		}
		this.head = node
		this.tail = node
	} else {
		node := &Node{
			val: value,
		}
		this.tail.Next, node.Pre, this.tail = node, this.tail, node
	}

	this.length++

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == nil {
		return false
	}

	if this.length > 1 {
		this.head, this.head.Next.Pre = this.head.Next, nil
	} else {
		this.head, this.tail = nil, nil
	}
	this.length--

	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.tail == nil {
		return false
	}

	if this.length > 1 {
		this.tail, this.tail.Pre.Next = this.tail.Pre, nil
	} else {
		this.head, this.tail = nil, nil
	}
	this.length--

	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.head == nil {
		return -1
	}

	return this.head.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.tail == nil {
		return -1
	}

	return this.tail.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.length == 0 {
		return true
	}

	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.length == this.capacity {
		return true
	}

	return false
}
