![img.png](images/img.png)

benchmark 的默认时间是 1s，那么我们可以使用 -benchtime 指定为 3s
![img_1.png](img_1.png)

-benchtime 的值除了是时间外，还可以是具体的执行次数。
![img_2.png](img_2.png)

-count 参数可以用来设置benchmark的轮数。
![img_3.png](img_3.png)

使用 -benchmem 参数看到内存分配的情况 Generate 分配的内存是 GenerateWithCap 的 6 倍，设置了切片容量，内存只分配一次，而不设置切片容量，内存分配了 40 次。 GenerateWithCap 的耗时比
Generate 少 20%。
![img_5.png](img_5.png)

复杂读观察 输入变为原来的 10 倍，函数每次调用的时长也差不多是原来的 10 倍 O(n)
![img_7.png](img_7.png)

通过b.ResetTimer() 将数据准备的代码耗时忽略掉。
![img_8.png](img_8.png)

通过 StopTimer & StartTimer 将数据准备/清理的代码耗时忽略掉。
![img_10.png](img_10.png)
![img_9.png](img_9.png)

