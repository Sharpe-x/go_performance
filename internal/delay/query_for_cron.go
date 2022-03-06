package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// 延时任务 和定时任务 的区别
// 定时任务有明确的触发时间 延时任务没有
// 定时任务有执行周期 延时任务在某事件触发后一段时间内执行 没有执行周期
// 定时任务一般执行的是批处理操作是多个任务 而延时任务一般是单个任务

// 延迟队列应该具备如下的特性
// 消息传输可靠 : 消息进入到延迟队列后 保证至少被消费一次 有消费失败的重试逻辑
// 高可用 + 可扩展性 支持多实例部署 挂掉一个实例后，还有备用实例继续提供服务
// 支持消息删除 业务使用方 可以随时删除指定消息
// 支持存储落地 任务支持延时/优先级和自动重试

// 方案一
// 1. 数据库轮询
// 通过一个线程定时给的去扫描数据库，通过订单时间来判断是否有超时的订单
// 优点 简单易行，支持集群操作
// 缺点 对服务器内存消耗大 存在延迟 取决于扫描频率 数据库损耗大

func delayByQueryDb() error {
	c := cron.New()

	_, err := c.AddFunc("@every 5s", func() {
		// query db find job
		fmt.Println("time now is: ", time.Now().String())
	})

	if err != nil {
		return errors.Wrap(err, "addFunc failed")
	}

	c.Start()
	return nil
}

func main() {
	err := delayByQueryDb()
	if err != nil {
		panic(err)
	}
	log.Println("start cron job")
	//select {}
}
