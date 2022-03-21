package designa_pattern

import "sync"

// 使用懒惰模式的单例没收 使用双重检查加锁保证线程安全

// Singleton 是单例模式
type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
