package main

import "fmt"

// SliceHeader 是切片的数据结构（运行时）
type SliceHeader struct {
	Data uintptr //  指针指向一块连续的内存空间
	Len  int
	Cap  int
}

func initSlice() {
	// 切片初始化方式
	// 1 通过下标的方式获得数组 或者切片的一部分 eg: arr[0:3] or  slice[0:3]
	arr := [3]int{1, 2, 3}
	slice := arr[0:3]
	fmt.Println(slice)
	arr[0] = 4
	// 使用下标初始化切片不会复制原数组或者原切片的大小和容量 所以修改新旧切片/数组的值 会修改原切片 的值
	fmt.Println("slice", slice) // slice [4 2 3]
	fmt.Println("arr", arr)     // arr [4 2 3]
	slice[1] = 5
	fmt.Println("slice", slice) //slice [4 5 3]
	fmt.Println("arr", arr)     // arr [4 5 3]

	slice = append(slice, 6, 7, 8) // 切片发生了扩容 cap 和 len 变为6 data 指向一块新的地址
	fmt.Println("slice", slice)    // slice [4 5 3 6 7 8]
	fmt.Println("arr", arr)        //arr [4 5 3]

	arr[2] = 6                  // 修改原数组 不会影响slice 因为底层数组的地址发生了变化
	fmt.Println("slice", slice) // slice [4 5 3 6 7 8]
	fmt.Println("arr", arr)     // arr [4 5 6]

	slice[3] = 9                // 同理 修改slice 不会影响原数组 因为slice底层数组的地址发生了变化
	fmt.Println("slice", slice) // slice [4 5 3 9 7 8]
	fmt.Println("arr", arr)     // arr [4 5 6]

	fmt.Println("======================= 通过下标的方式获得数组 或者切片的一部分end =======================")
	// 2 使用字面量创建 slice eg ；[]int{1,2,3}
	// 2.1 根据切片中的元素数量推断底层数组的大小并创建一个数组
	// 2.2 将这些字面量元素存储到初始化的数组中
	// 2.3 创建一个同样指向【3]int 类型的指针
	// 2.4 将静态存储区的数组赋值给2.3 创建的指针
	// 2.5 通过[:] (使用下标创建切片的方法) 获取一个地岑使用2.1 创建的数组 的切片
	a := []int{1, 2, 3, 4, 5} // 定义并初始化
	fmt.Println(a)
	fmt.Println("======================= 通过字面量创建 end =======================")

	// 3 使用关键字 make 创建slice
	// 3.1 检查 会检查 传入的容量cap 一定大于等于len
	// 3.2 会检查切片大小和容量是不是足够小 如果太大 会发生逃逸 最终在堆中初始化 大于32kb 的对象会在堆中初始化
	b := make([]int, 3, 4)
	fmt.Println("len(b) = ", len(b), ",cap(b) = ", cap(b))
	fmt.Println("======================= 使用关键字创建 end =======================")
}

func useSlice() {
	// append
	slice := []int{1, 2, 3, 4, 5}
	newSlice := append(slice, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(newSlice)

	// 初始化slice1 已经有3个值了 0 0 0
	slice1 := make([]int, 3, 6)
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(slice1)                                                    // 0 0 0 1 2 3
	fmt.Println("len(slice1) = ", len(slice1), ",cap(slice)", cap(slice1)) // 6,6

	newSlice1 := append(slice1, 4, 5, 6)
	fmt.Println("len(newSlice1) = ", len(newSlice1), ",cap(slice)", cap(newSlice1)) // 发生了扩容(容量翻倍 具体规则见下面)  9,12
	fmt.Println(newSlice1)                                                          //  0 0 0 1 2 3 4 5 6

	fmt.Println("modify slice1[1]")
	slice1[1] = -1
	fmt.Println(slice1)    // 0 -1 0 1 2 3
	fmt.Println(newSlice1) // 0 0 0 1 2 3 4 5 6 // 已经是新的数组了 所以不影响

	fmt.Println("======================= slice1  newSlice1 end =======================")
	slice2 := make([]int, 3, 6)
	slice2 = append(slice2, 1, 2, 3, 4)                                     //会扩容
	fmt.Println(slice2)                                                     // 0 0 0 1 2 3 4
	fmt.Println("len(slice2) = ", len(slice2), ",cap(slice2)", cap(slice2)) // 7,12

	newSlice2 := append(slice2, 5, 6, 7)                                            // 没有扩容 现在只有10 个元素
	fmt.Println("len(newSlice2) = ", len(newSlice2), ",cap(slice)", cap(newSlice2)) // 10,12
	fmt.Println(newSlice2)                                                          //  0 0 0 1 2 3 4 5 6 7

	fmt.Println("modify slice1[1]")
	slice2[1] = -1
	fmt.Println(slice2)    // 0 -1 0 1 2 3 4
	fmt.Println(newSlice2) // 0 -1 0 1 2 3 4 5 6 7// 和slice1 共用一段数组 所以会受影响

	slice2 = append(slice2, 9, 9, 9, 9, 9)                                  // 没有触发扩容 所以会影响 newSlice2 假如触发了扩容 就不会影响 newSlice2（赋值发生在扩容之后,赋值是在新的地址上）
	fmt.Println("len(slice2) = ", len(slice2), ",cap(slice2)", cap(slice2)) // 12,12
	fmt.Println(slice2)                                                     // 0 -1 0 1 2 3 4 9 9 9 9 9
	fmt.Println(newSlice2)                                                  //0 -1 0 1 2 3 4 9 9 9 / 和slice2 共用一段数组 所以会受影响

	fmt.Println("======================= slice2  newSlice2 end =======================")

	// 新切片容量的确定策略
	// 1 如果期望容量大于 当前容量的2 倍 就会使用期望容量
	// 2 如果当前切片的长度小于1024 就将容量翻倍
	// 3 如果当前切片的长度大于1024 每次增加25% 直到大于期望容量
	// 4 1,2,3 只是大致确定了切片的大致容量  还要根据切片中的元素大小对齐内存 向上取整 (减少内存碎片)
}

func copySlice() {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{5, 6}
	copy(sliceA, sliceB)
	fmt.Println("len(sliceA) = ", len(sliceA), ",cap(sliceA) = ", cap(sliceA))
	fmt.Println("len(sliceB) = ", len(sliceB), ",cap(sliceB) = ", cap(sliceB))
	// copy( destSlice, srcSlice []T) int
	// 其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice）
	//，目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。

	fmt.Println("======================= copySlice =======================")
}

func sliceAsParam(slice []string) {
	slice = append(slice, "append", "Inner")
	fmt.Println("in sliceAsParam :", slice)
	fmt.Println("len(slice) = ", len(slice), "cap(slice)  = ", cap(slice))
}

func sliceAsParam2(slice []string) {
	slice[0] = "sliceAsParam2"
}

func sliceAsParam3(slice []string) {
	slice = append(slice, "append", "Inner")
	fmt.Println("in sliceAsParam :", slice)                                // in sliceAsParam : [hello this is main append Inner]
	fmt.Println("len(slice) = ", len(slice), "cap(slice)  = ", cap(slice)) // len(slice) =  6 cap(slice)  =  8
	slice[0] = "sliceAsParam2"
}

func main() {
	initSlice()
	useSlice()
	copySlice()
	slice := []string{
		"Hello", "this", "is", "main",
	}
	// 值传递 外部不受影响
	sliceAsParam(slice)              // in sliceAsParam : [Hello this is main append Inner]
	fmt.Println("out slice:", slice) // out slice: [Hello this is main]
	fmt.Println("len(slice) = ", len(slice), "cap(slice)  = ", cap(slice))
	fmt.Println("=========================扩容======================")

	sliceNotAddCap := make([]string, 4, 6)
	sliceNotAddCap[0] = "Hello"
	sliceNotAddCap[1] = "this"
	sliceNotAddCap[2] = "is"
	sliceNotAddCap[3] = "main"
	sliceAsParam(sliceNotAddCap)                                                                               // in sliceAsParam : [Hello this is main append Inner] len(slice) =  6 cap(slice)  =  6
	fmt.Println("out sliceNotAddCap:", sliceNotAddCap)                                                         // out sliceNotAddCap: [Hello this is main]
	fmt.Println("len(sliceNotAddCap) = ", len(sliceNotAddCap), "cap(sliceNotAddCap)  = ", cap(sliceNotAddCap)) // len(sliceNotAddCap) =  4 cap(sliceNotAddCap)  =  6
	fmt.Println("=========================扩容======================")

	sliceStr := []string{"hello", "this", "is", "main"}
	sliceAsParam2(sliceStr)
	fmt.Println("out sliceStr:", sliceStr) // out sliceStr: [sliceAsParam2 this is main] 受了影响
	fmt.Println("len(sliceStr) = ", len(sliceStr), "cap(sliceStr)  = ", cap(sliceStr))

	fmt.Println("=========================扩容======================")

	sliceStr2 := []string{"hello", "this", "is", "main"}
	sliceAsParam3(sliceStr2)                                                               // 扩容之后才append 不影响外部slice
	fmt.Println("out sliceStr2:", sliceStr2)                                               // out sliceStr2: [hello this is main]
	fmt.Println("len(sliceStr2) = ", len(sliceStr2), "cap(sliceStr2)  = ", cap(sliceStr2)) // len(sliceStr2) =  4 cap(sliceStr2)  =  4

}
