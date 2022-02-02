package concurrency

import (
	"runtime"
	"testing"
	"time"
)

func Test_useChannelControlGoroutineNum(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	useChannelControlGoroutineNum()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
