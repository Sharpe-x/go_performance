package concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_timeout(t *testing.T) {
	fmt.Println(timeout(doBadThings))
}

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		fmt.Println(timeout(f))
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestBadThings(t *testing.T) {
	test(t, doBadThings)
}

func TestBufferTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(timeoutWithBuffer(doBadThings))
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestGoodTimeout(t *testing.T) {
	test(t, doGoodThings)
}
