package forrange

import (
	"fmt"
	"math/rand"
	"time"
)

// range 可以用来很方便地遍历数组(array)、切片(slice)、字典(map)和信道(chan)
//range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，那么 for 和 range 的性能几乎是一样，例如 []int。但是如果迭代的元素内存占用较高，例如一个包含很多属性的 struct
//结构体，那么 for 的性能将显著地高于 range，有时候甚至会有上千倍的性能差异。对于这种场景，建议使用 for，如果使用 range，建议只迭代下标，通过下标访问迭代值，这种使用方式和 for
//就没有区别了。如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。

func rangeExample() {
	//变量 strs 在循环开始前，仅会计算一次，如果在循环中修改切片的长度不会改变本次循环的次数。
	// 迭代过程中，每次迭代的下标和值被赋值给变量 i 和 str，第二个参数 str 是可选的。
	// 针对 nil 切片，迭代次数为 0。
	strs := []string{"Go", "Rust", "Java", "Python", "C++"}
	for i, str := range strs {
		//strs = append(strs, str)
		fmt.Println(i, str)
	}
	//fmt.Println(strs)

	// for 没什么差异
	for i := range strs {
		fmt.Println(i, strs[i])
	}

	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range testMap {
		// 迭代过程中，删除还未迭代到的键值对，则该键值对不会被迭代。
		delete(testMap, "two")
		// 迭代过程中，如果创建新的键值对，那么新增键值对，可能被迭代，也可能不会被迭代。
		// 针对 nil 字典，迭代次数为 0
		testMap["four"] = 4
		fmt.Println(k, v)
	}

	ch := make(chan string)
	go func() {
		ch <- "Go"
		ch <- "Test"
		ch <- "Bench"
		ch <- "Benchmark"
		close(ch)
	}()

	// 发送给信道(channel) 的值可以使用 for 循环迭代，直到信道被关闭。
	// 如果是 nil 信道，循环将永远阻塞。
	for str := range ch {
		fmt.Println(str)
	}

}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generateItems(n int) []*Item {
	items := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item{
			id: i,
		})
	}
	return items
}

type Item struct {
	id   int
	name string
	val  [4096]byte
}
