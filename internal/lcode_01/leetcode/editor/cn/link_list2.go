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

func reverseListRecByK(head *Node, k int) *Node {
	dummyHead := &Node{
		Val:  -1,
		Next: head,
	}
	pre, cur := dummyHead, head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	cur = head

	for i := 0; i < length/k; i++ {
		for j := 1; j < k; j++ {
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		pre = cur
		cur = pre.Next
	}
	return dummyHead.Next
}

func myReorderList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var nodes []*Node
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	i, j := 0, len(nodes)-1
	for i != j {
		nodes[i].Next = nodes[j]
		i++

		if i != j {
			nodes[j].Next = nodes[i]
			j--
		}
	}

	nodes[i].Next = nil
	return head
}

func main() {
	/*a := []int{1, 2, 3, 4, 5}
	list := newList(a)
	listPrint(list)

	sList := listSwapPairs(list)
	listPrint(sList)*/

	/*rList := listReverse(list)
	listPrint(rList)*/

	/*b := []int{1, 2, 3, 4, 5}
	bList := newList(b)
	resList := reverseListRecByK(bList, 3)
	listPrint(resList)*/

	c := []int{1, 2, 3, 4, 5}
	cList := newList(c)
	listPrint(cList)
	recorderList := myReorderList(cList)
	listPrint(recorderList)
}
