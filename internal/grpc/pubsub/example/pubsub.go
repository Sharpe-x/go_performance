package main

import (
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.Contains(key, "golang") {
				return true
			}
		}
		return false
	})

	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.Contains(key, "docker") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang java")
	go p.Publish("docker k8s")

	time.Sleep(1)
	go func() {
		fmt.Println("golang topic", <-golang)
	}()

	go func() {
		fmt.Println("docker topic", <-docker)
	}()

	time.Sleep(time.Second * 10)
}
