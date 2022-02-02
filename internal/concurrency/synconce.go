package concurrency

import (
	"log"
	"os"
	"strconv"
	"sync"
)

// sync.Once 是Go 标准库提供的使函数只执行一次的实现 常应用于单例模式 例如初始化配置、保持数据库连接等 作用于init 函数类似 但有区别
// init 函数是当所在的package 首次被加载时执行 若迟迟未被使用 即浪费了内存 又延长了程序加载时间
// sync.Once 可以在代码的任意位置初始化和调用 因此可以延迟道使用时再执行 并发场景下是线程安全的

//sync.Once 被用于控制变量的初始化 这个变量的读写满足如下3个条件：
/* 当且仅当第一次访问某个变量时 进行初始化写
 变量初始化过程中 所有读都被阻塞 直到初始化完成
变量仅初始化一次 初始化完成后驻留在内存里*/

//sync.Once 仅提供了一个方法 Do 参数f 是对象初始化函数
// func (o *Once) Do (f func())

//函数 ReadConfig 需要读取环境变量，并转换为对应的配置。环境变量在程序执行前已经确定，执行过程中不会发生改变。ReadConfig 可能会被多个协程并发调用，为了提升性能（减少执行时间和内存占用），
//使用 sync.Once 是一个比较好的方式。

type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{
			Server: os.Getenv("TT_SERVER_URL"),
		}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		if err != nil {
			config.Port = 8000
		}
		log.Println("init config")
	})
	return config
}
