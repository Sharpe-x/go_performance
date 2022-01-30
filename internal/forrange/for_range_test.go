package forrange

import "testing"

func Test_rangeExample(t *testing.T) {
	rangeExample()
}

// 遍历 []int 类型的切片，for 与 range 性能几乎没有区别。
func BenchmarkForIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		length := len(nums)
		var tmp int
		for j := 0; j < length; j++ {
			tmp = nums[j]
		}
		_ = tmp
	}
}

func BenchmarkRangeIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range nums {
			tmp = num
		}
		_ = tmp
	}
}

//=========================================
// 结论
// 仅遍历Item 数组下标的情况下，for 和 range 的性能几乎是一样的。
// for 的性能大约是 range (同时遍历下标和值) 上千倍

// 原因
// range 对每个迭代值都创建了一个拷贝。因此如果每次迭代的值内存占用很小的情况下，for 和 range 的性能几乎没有差异，
//但是如果每个迭代值内存占用很大，例如上面的例子中，每个结构体需要占据 4KB+ 的内存，这种情况下差距就非常明显了。

// 使用 for 迭代时 将修改每个结构体，有效。
// 使用 range 迭代（Item 数组）时，修改每个结构体无效，因为 range 返回的是拷贝。
func BenchmarkForStructSlices(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for j := 0; j < length; j++ {
			tmp = items[j].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStructSlices(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := range items {
			tmp = items[j].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStructSlices(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

//=========================================
// 切片元素从结构体 Item 替换为指针 *Item 后，for 和 range 的性能几乎是一样的。
//而且使用指针还有另一个好处，可以直接修改指针对应的结构体的值。

//
func BenchmarkRangePointerSlices(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

func BenchmarkForPointerSlices(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for j := 0; j < length; j++ {
			tmp = items[j].id
		}
		_ = tmp
	}
}
