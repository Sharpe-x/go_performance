package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}         //订阅者为一个通道
	topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// Publisher 发布者对象
type Publisher struct {
	m           sync.RWMutex
	buffer      int //订阅队列的缓存大小
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

// NewPublisher 构建一个发布者 设置缓存大小和发布超时时间
func NewPublisher(publisherTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publisherTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加一个订阅者 订阅所有主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 订阅一个主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Evict 退出主题
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// Publish 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// sendTopic 发送主题
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// Close 关闭发布对象 和所有订阅者通通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello world")
	p.Publish("hello golang")

	go func() {
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()

	time.Sleep(time.Second * 5)
}
