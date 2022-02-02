package concurrency

import (
	"runtime"
	"testing"
	"time"
)

func Test_sendTasks(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasks()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}

func Test_sendTaskCheckClose(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTaskCheckClose()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}

func Test_testMyChannel(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	testMyChannel()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
