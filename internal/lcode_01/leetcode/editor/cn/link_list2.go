package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func listPrint(head *Node) {

	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println("NULL")
}

func newList(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	head := &Node{
		Val:  arr[0],
		Next: nil,
	}

	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &Node{
			Val:  arr[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return head
}

func listReverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	var pre *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func listSwapPairs(head *Node) *Node {

	dummyHead := &Node{
		Val:  -1,
		Next: head,
	}

	pre := dummyHead
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := node1.Next
		next := node2.Next

		node2.Next = node1
		node1.Next = next

		pre.Next = node2
		pre = node1
	}

	return dummyHead.Next
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	list := newList(a)
	listPrint(list)

	sList := listSwapPairs(list)
	listPrint(sList)

	/*rList := listReverse(list)
	listPrint(rList)*/
}
