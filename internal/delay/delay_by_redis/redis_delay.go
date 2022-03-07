package delay_by_redis

// 利用redis的zset,zset是一个有序集合，每一个元素(member)都关联了一个score,通过score排序来取集合中的值。
// 通过 ZRANGEBYSCORE ，可以得到基于 Score 在指定区间内的元素（排序）。基于 Sorted Set 的延时队列模型如下：
//
//SortSet 的 key 作为业务维度的属性（队列）名字，比如一种命名方式为 <业务: 命名空间: 队列名>
//SortSet 中的元素做为任务消息，Score 视为本任务延迟的时间（戳）
