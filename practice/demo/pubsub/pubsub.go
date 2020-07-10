package main

import (
	"fmt"
	"sync"
	"time"
	"strings"
)

// 发布者/订阅者模型
type (
	subscriber chan interface{}        // 订阅者
	topicFunc func(v interface{}) bool // 订阅主题过滤器函数
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列缓存大小
	timeout     time.Duration            // 发布超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// 构建一个发布者, 可以设置发布时间和缓存队列长度
func NewPublisher(buffer int, publishTimeout time.Duration) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个订阅者, 订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 添加一个订阅者, 订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	sub := make(chan interface{}, p.buffer)
	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[sub] = topic

	return sub
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	fmt.Println("Start Publish ", v)
	p.m.RLock()
	defer p.m.RUnlock()

	wg := sync.WaitGroup{}
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

	fmt.Println("Start Wait...")
	wg.Wait()
	fmt.Println("End Wait...")
}

// 发布一个主题, 可以容忍一定超时
func (p *Publisher) sendTopic(sub chan interface{}, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Start sendTopic, ", v)
	if topic != nil && !topic(v) {
		fmt.Println("topic suc ", v)
		return
	}

	select {
	case sub <- v:
		fmt.Printf("Send suc, sub %v, v %v\n", sub, v)
	case <- time.After(p.timeout):
		fmt.Printf("Send timeout, sub %v, v %v\n", sub, v)
	}
}

// 关闭发布者对象, 同时关闭所有的订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}


func main() {
	p := NewPublisher(3, 100 * time.Millisecond)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}

		return false
	})

	p.Publish("hello, world")
	p.Publish("hello, golang")

	go func() {
		for msg := range all {
			fmt.Println("all RECEIVED: ", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang RECEIVED: ", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}