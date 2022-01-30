package slice

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func learnSlices() {
	a := []int32{1, 2, 3, 4, 5}
	b0 := make([]int32, len(a))
	copy(b0, a) //  切片的copy
	fmt.Printf("copy a to b0, a = %v, b0 = %v\n", a, b0)

	b1 := make([]int32, len(a))
	b1 = append([]int32(nil), a...)
	fmt.Printf("copy a to b1, a = %v, b1 = %v\n", a, b1)

	b2 := make([]int32, len(a))
	b2 = append(a[:0:0], a...)
	fmt.Printf("copy a to b2, a = %v, b2 = %v\n", a, b2)

	//切片有三个属性，指针(ptr)、长度(len) 和容量(cap)。append 时有两种场景：
	// 当 append 之后的长度小于等于 cap，将会直接利用原底层数组剩余的空间。
	// 当 append 后的长度大于 cap 时，则会分配一块更大的区域来容纳新的底层数组。
	// 为了避免内存发生拷贝，如果能够知道最终的切片的大小，预先设置 cap 的值能够获得最好的性能。
	a = append(a, b0...)
	fmt.Printf("append b0, a = %v, b0 = %v\n", a, b0)

	//切片的底层是数组，因此删除意味着后面的元素需要逐个向前移位。
	//每次删除的复杂度为 O(N)，因此切片不合适大量随机删除的场景，这种场景下适合使用链表。
	a = append(a[:5], a[6:]...)
	fmt.Printf("delete a[5]=, after a = %v\n", a)

	a = a[:5+copy(a[5:], a[5+1:])]
	fmt.Printf("delete a[5]=, after a = %v\n", a)

	// insert 和 append 类似。即在某个位置添加一个元素后，将该位置后面的元素再 append 回去。复杂度为 O(N)。
	//不适合大量随机插入的场景。
	a = append(a[:5], append([]int32{1, 2}, a[5:]...)...)
	fmt.Printf("append 1,2 =, after a = %v\n", a)

	// 在末尾追加元素，不考虑内存拷贝的情况，复杂度为 O(1)。
	a = append(a, 6)
	fmt.Printf("append 6, after a = %v\n", a)

	// 在头部追加元素，时间和空间复杂度均为 O(N)，不推荐。
	a = append([]int32{0}, a...)
	fmt.Printf("append 0, after a = %v\n", a)

	// 尾部删除元素，复杂度 O(1)
	deleteNum, a := a[len(a)-1], a[:len(a)-1]
	fmt.Printf("deleteNum = %v,a = %v\n", deleteNum, a)

	// 头部删除元素，如果使用切片方式，复杂度为 O(1)。但是需要注意的是，底层数组没有发生改变，
	//第 0 个位置的内存仍旧没有释放。如果有大量这样的操作，头部的内存会一直被占用。
	deleteNum, a = a[0], a[1:]
	fmt.Printf("deleteNum = %v,a = %v\n", deleteNum, a)

	// 在已有切片的基础上进行切片，不会创建新的底层数组。因为原来的底层数组没有发生变化，
	//内存会一直占用，直到没有变量引用该数组。因此很可能出现这么一种情况，原切片由大量的元素构成，
	//但是我们在原切片的基础上切片，虽然只使用了很小一段，但底层数组在内存中仍然占据了大量空间，得不到释放。
	// 推荐的做法，使用 copy 替代 re-slice。 func lastNumsByCopy

}

// 直接在原切片基础上进行切片。
func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

// origin 的最后两个元素拷贝到新切片上，然后返回新切片。
func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

// generateWithCap 用于随机生成 n 个 int 整数，64位机器上，一个 int 占 8 Byte，128 * 1024 个整数恰好占据 1 MB 的空间。
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// printMem 于打印程序运行时占用的内存大小。
func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}
