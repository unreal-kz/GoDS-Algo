package main

import "fmt"

type linkedList struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}

//insert
func (l *linkedList) insert(n *node) {
	currNode := l.head
	l.head = n
	l.head.next = currNode
	l.length++
}

//delete
func (l *linkedList) delete(v int) {
	if l.length == 0 {
		return
	}
	if l.head.data == v {
		l.head = l.head.next
		l.length--
		return
	}

	prevNode := l.head
	for prevNode.next.data != v {
		if prevNode.next.next == nil {
			return
		}
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
	l.length--

}

//custom print function
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func main() {
	mList := &linkedList{}
	node1 := &node{data: 42}
	node2 := &node{data: 43}
	node3 := &node{data: 45}
	node4 := &node{data: 15}
	node5 := &node{data: 25}
	node6 := &node{data: 5}

	mList.insert(node1)
	mList.insert(node2)
	mList.insert(node3)
	mList.insert(node4)
	mList.insert(node5)
	mList.insert(node6)

	mList.printListData()
	mList.delete(25)
	mList.printListData()
	mList.delete(5)
	mList.printListData()
	mList.delete(125)
	mList.printListData()
}
