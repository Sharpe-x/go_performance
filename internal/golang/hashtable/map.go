package main

import "fmt"

// 哈希函数能将不同键映射到不同索引上 这要求哈希函数的输出范围大于输入,但由于键的数量远远大于映射范围，所以效果不是很理想
// 比较实际的方式是让哈希函数的结果尽可能的均匀分布，然后通过工程手段解决哈希冲突问题。
// 哈希冲突不是多个键对应的哈希完全相等，可能是多个哈希的部分相等，例如2个键对应哈希的前4个字节相同

// 解决冲突的方法
// - 开放寻址法
// 核心思想 依次探测和比较数组中的元素以判断目标键值队是否存在于哈希表中
// 根据hash值对数组取摸得到索引 然后从索引开始插入（如果有值索引加1） 或者 开始遍历
// 装载因子 数组中元素数量和数组大小的比值 装载因子越大 性能越差

// - 拉链法
// 实现拉链一般是数组加上链表（也有可能会引入红黑树来提高性能），拉链法会使用链表数组作为哈希底层的数据结构 可以看成是可拓展的二维数组
// 此时的index 代表的就是桶 然后在桶中寻找键值对
// 装载因子 元素数量/桶数量 装载因子越大 性能越差

// Go 的实现

type hmap struct {
	count      int // 哈希表中的数量
	flags      uint8
	B          uint8 // 当前哈希表持有的buckets 数量
	noverflow  uint16
	hash0      uint32  // 哈希表的种子 为哈希函数引入随机性
	oldbuckets uintptr // 哈希表扩容时保存之前的buckets的字段
	extra      *mapextra
}
type mapextra struct {
	overflow     *[]*bmap
	oldoverflow  *[]*bmap
	nextOverflow *bmap
}
type bmap struct {
	/*	tophash [buckentCnt]uint8*/
}

func initMap() {
	// 字面量初始化
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	//当 哈希表中的数量少于或者等于25个时，编译器会将字面量初始化的结构体优化为以下代码
	newHash := make(map[string]int, 3)
	hash["1"] = 1
	hash["2"] = 2
	hash["3"] = 3

	// 一旦哈希表中的元素超过25 个 编译器就会创建2个数组来分别存储键和值 然后通过for 循环假如哈希表

	fmt.Println(hash, newHash)

	// 当桶的数量多于24 个时，会额外创建2 的B-4 次方个溢出桶

	// 哈希表扩容
	// - 装载因子大于6.5 时
	// - 哈希表使用了太多溢出桶
	// 扩容不是原子过程
	// 哈希表在存储元素过多时会触发扩容操作，每次都会将桶的数量翻倍，扩容不是原子的，而是通过time.growWork 增量触发的。在扩容期间访问哈希表会使用旧桶，向哈希表写入数据时会触发
	// 旧桶元素的分流。除了这种正常扩容 为了解决大量写入、删除造成的内存泄漏问题，哈希表引入了sameSizeGrow 机制，在出现较多溢出桶时会整理哈希表的内存来减少空间占用

}

func test(m map[int]string) {
	fmt.Printf("%p \n", m)
	m[1] = "test modify"
	m[10] = "test add"
}

func test2(m2 map[int]string) {
	m2 = make(map[int]string, 2)
	m2[1] = "test2 modify"
	m2[10] = "test2 add"
}

func main() {
	initMap()

	map1 := map[int]string{
		1: "hi",
		2: "hell",
	}
	fmt.Printf("%p \n", map1)

	test(map1)
	fmt.Println(map1)

	var map2 map[int]string // 地址为初始化 和分配
	fmt.Printf("%p \n", map2)
	test2(map2) // 内部初始化了 但是是值传递 所以不影响外部
	fmt.Println(map2)
	map2 = make(map[int]string, 2)
	map2[1] = "main modify"
	map2[10] = "main add"
	fmt.Printf("%p \n", map2)
	fmt.Println(map2)
}
