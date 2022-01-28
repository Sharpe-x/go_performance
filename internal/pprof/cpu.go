package pprof

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

/*CPU 性能分析(CPU profiling) 是最常见的性能分析类型。
启动 CPU 分析时，运行时(runtime) 将每隔 10ms 中断一次，记录此时正在运行的协程(goroutines) 的堆栈信息。
程序运行结束后，可以分析记录的数据找到最热代码路径(hottest code paths)。
一个函数在性能分析数据中出现的次数越多，说明执行该函数的代码路径(code path)花费的时间占总运行时间的比重越大。*/
// Go 的运行时性能分析接口都位于 runtime/pprof 包中。只需要调用 runtime/pprof 库即可得到我们想要的数据。

// https://geektutu.com/post/hpg-pprof.html#2-CPU-%E6%80%A7%E8%83%BD%E5%88%86%E6%9E%90

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}

	}
	return nums
}

func GetCpuProf() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic("close file failed: " + err.Error())
		}
	}(f)
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}
