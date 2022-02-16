package main

import (
	"context"
	"fmt"
)

type processHandler func(ctx context.Context, str string) error
type myInterceptor func(ctx context.Context, str string, h processHandler) error

func chainUnaryInterceptor(interceptors ...myInterceptor) myInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, str string, h processHandler) error {
		chainer := func(currentInter myInterceptor, currentHandler processHandler) processHandler {
			return func(ctx context.Context, str string) error {
				return currentInter(ctx, str, currentHandler)
			}
		}

		chainedHandler := h
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, str)
	}
}

func testMyInterceptor() {
	pHandler := func(ctx context.Context, str string) error {
		fmt.Println(str)
		return nil
	}

	var inter1 = func(ctx context.Context, str string, h processHandler) error {
		fmt.Println("inter1 before")
		err := h(ctx, str)
		fmt.Println("inter1 after")
		return err
	}

	var inter2 = func(ctx context.Context, str string, h processHandler) error {
		fmt.Println("inter2 before")
		err := h(ctx, str)
		fmt.Println("inter2 after")
		return err
	}

	var inter3 = func(ctx context.Context, str string, h processHandler) error {
		fmt.Println("inter3 before")
		return h(ctx, str)
	}

	interOne := chainUnaryInterceptor(inter1, inter2, inter3)
	_ = interOne(context.Background(), "测试一哈哈", pHandler)
}
