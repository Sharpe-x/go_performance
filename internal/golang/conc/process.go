package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/conc/stream"
	"runtime/debug"
	"time"
)

type caughtPanicError struct {
	val   any
	stack []byte
}

func (e *caughtPanicError) Error() string {
	return fmt.Sprintf(
		"panic: %q\n%s",
		e.val,
		string(e.stack),
	)
}

func doSomeThingsPanic() {
	time.Sleep(3 * time.Second)
	panic("do something panic")
}

func main() {
	//normalRun()
	//conRun()

	/*srcChan := make(chan int)

	go func() {
		process(srcChan)
	}()

	for i := 0; i < 10; i++ {
		srcChan <- i
	}
	close(srcChan)

	time.Sleep(time.Minute * 2)*/

	//processWithResult()

	ExampleStream()
}

func conRun() {
	var wg conc.WaitGroup
	wg.Go(doSomeThingsPanic)
	wg.Wait()
}

func normalRun() {
	done := make(chan error)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				done <- &caughtPanicError{
					val:   err,
					stack: debug.Stack(),
				}
			} else {
				done <- nil
			}
		}()
		doSomeThingsPanic()
	}()

	err := <-done
	if err != nil {
		panic(err)
	}
}

func process(stream chan int) {
	p := pool.New().WithMaxGoroutines(3)
	p.WithErrors()
	for elem := range stream {
		e := elem
		p.Go(func() {
			handle(e)
		})
	}
	p.Wait()
}

func processWithResult() {
	p := pool.NewWithResults[int]()
	for i := 0; i < 10; i++ {
		num := i
		p.Go(func() int {
			return num * 2
		})
	}
	res := p.Wait()
	fmt.Println(res)
}

func handle(e int) {
	fmt.Printf("handle %d start\n", e)
	time.Sleep(time.Duration(e) * time.Second)
	if e%2 == 0 {
		panic(fmt.Sprintf("panic for test %d", e))
	}
	fmt.Printf("handle %d end\n", e)
}

func ExampleStream() {
	times := []int{20, 52, 16, 45, 4, 80}

	s := stream.New()
	for _, millis := range times {
		dur := time.Duration(millis) * time.Millisecond
		s.Go(func() stream.Callback {
			time.Sleep(dur)
			// This will print in the order the tasks were submitted
			return func() { fmt.Println(dur) }
		})
	}
	s.Wait()

	// Output:
	// 20ms
	// 52ms
	// 16ms
	// 45ms
	// 4ms
	// 80ms
}
