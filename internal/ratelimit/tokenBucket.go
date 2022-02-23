package main

import (
	"fmt"
	"time"
)

const fillInterval = time.Millisecond * 10

var tokenBucket = make(chan struct{}, 100)

func fillToken() {
	ticker := time.NewTicker(fillInterval)
	for {
		select {
		case <-ticker.C:
			select {
			case tokenBucket <- struct{}{}:
			default:
			}
			fmt.Println("current token cnt:", len(tokenBucket), time.Now())
		}
	}
}

func TakeAvailableToken(block bool) bool {
	var tokenResult bool
	if block {
		select {
		case <-tokenBucket:
			tokenResult = true
		}
	} else {
		select {
		case <-tokenBucket:
			tokenResult = true
		default:
			tokenResult = false
		}
	}
	return tokenResult
}

func main() {
	go fillToken()
	time.Sleep(time.Hour)
}

// 上一次放令牌的时间为t1 当时的令牌数为k1 放令牌的时间间隔为ti 令牌桶的容量为cap
// t2 时刻来取令牌 此时令牌桶中理论有多少个令牌呢
// cur = k1 + ((t2 -t1)/ti) * x
// cur = cur > cap ? cap : cur
