package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		cur = cur.Next
	}

	return head
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " -> ")
		cur = cur.Next
	}
	fmt.Println("NULL")
}

func reverseListUsePre(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverseListRec(head *ListNode) *ListNode {
	// 递归终止条件
	if head == nil || head.Next == nil {
		return head
	}

	rHead := reverseListRec(head.Next)
	// head.Next 此刻指向head后面链表的尾节点
	// head->next->next = head 把head节点放在了尾部
	head.Next.Next = head
	head.Next = nil

	return rHead
}

func removeElementsInList(head *ListNode, val int) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	cur := dummy

	for cur.Next != nil {
		if cur.Next.Val == val {
			del := cur.Next
			cur.Next = del.Next
			del.Next = nil
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

func mySwapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	p := dummyHead

	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next

		node2.Next = node1
		node1.Next = next
		p.Next = node2
		p = node1
	}

	return dummyHead.Next
}

func reverseGroupByK(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	pre := dummy
	cur := head
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	head = dummy.Next
	for i := 0; i < length/k; i++ {
		for j := 0; j < k-1; j++ {
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		pre = cur
		cur = pre.Next
	}
	return dummy.Next
}

func mydeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{
		Val:  -999,
		Next: head,
	}
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

func main() {

	/*arr := []int{1, 2, 3, 4, 5}
	head := createList(arr)
	printList(head)

	head2 := reverseListUsePre(head)
	printList(head2)

	head3 := reverseListRec(head2)
	printList(head3)

	arr = []int{1, 2, 3, 4, 5, 6}
	head = reverseListRec(createList(arr))
	printList(head)

	arr = []int{1, 2, 3, 4, 5, 6}
	head = createList(arr)
	printList(head)
	head = removeElementsInList(head, 1)
	printList(head)

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	head = createList(arr)
	printList(mySwapPairs(head))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	bHead := createList(arr)
	printList(reverseGroupByK(bHead, 3))*/

	arr := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	list := createList(arr)
	dHead := deleteDuplicates(list)
	printList(dHead)
}
