package designa_pattern

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)

	instances := [100]*Singleton{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < 100; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
