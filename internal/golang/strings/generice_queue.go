package main

import "fmt"

type Queue[T interface{}] struct {
	elements []T
}

func (q *Queue[T]) Put(value T) {
	q.elements = append(q.elements, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, true
	}

	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func main() {
	testQueue()
}

func testQueue() {
	var q Queue[int]
	q.Put(1)
	q.Put(2)
	q.Put(3)

	fmt.Println(q.Size())
	value, ok := q.Pop()
	fmt.Println(value, ok)

	value, ok = q.Pop()
	fmt.Println(value, ok)

	value, ok = q.Pop()
	fmt.Println(value, ok)

	value, ok = q.Pop()
	fmt.Println(value, ok)
}
