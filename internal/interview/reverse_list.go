package main

import "fmt"

// 反转链表

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	for head != nil {
		node := head.next
		head.next = newHead
		newHead = head
		head = node
	}
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	ret := reverseList2(head.next)
	head.next.next = head
	head.next = nil
	return ret
}

func newList(n int) *ListNode {
	if n <= 0 {
		return nil
	}

	head := &ListNode{
		val:  1,
		next: nil,
	}

	cur := head
	for i := 1; i < n; i++ {
		node := &ListNode{
			val:  i + 1,
			next: nil,
		}
		cur.next = node
		cur = node
	}

	return head
}

func printList(list *ListNode) {
	if list == nil {
		return
	}

	cur := list
	for cur != nil {
		if cur.next == nil {
			fmt.Print(cur.val, "->null", "\n")
			return
		}
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
}

func main() {
	list := newList(5)
	printList(list)
	revList := reverseList(list)
	printList(revList)
}

// 给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，
// 并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）
