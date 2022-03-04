package concurrency

import (
	"fmt"
	"sync"
)

type task struct {
	i *int
	sync.WaitGroup
}

func Order() {
	works := make(chan *task, 10)
	fifo := make(chan *task, 10)

	for i := 1; i <= 10; i++ {
		num := i
		t := &task{
			i: &num,
		}
		sendToWorker(works, t)
		sendToKeepOrder(fifo, t)
	}

	close(works)
	close(fifo)

	worker(works)

	index := 0
	for tk := range fifo {
		tk.Wait()
		if index == 0 {
			fmt.Println()
			index++
		}
		fmt.Print(*tk.i, "==")
	}
}

func sendToKeepOrder(fifo chan *task, t *task) {
	fifo <- t
}

func worker(works chan *task) {

	for t := range works {
		go func(t *task) {
			defer t.Done()
			*t.i *= 10
			//fmt.Println("*t.i = ", *t.i)
		}(t)
	}
}

func sendToWorker(works chan *task, t *task) {
	t.Add(1)
	works <- t
}
