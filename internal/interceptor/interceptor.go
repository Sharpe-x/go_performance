package main

import (
	"context"
	"fmt"
)

// 假设有一个方法 handler(ctx context.Context) ，
// 想要给这个方法赋予一个能力：允许在这个方法执行之前能够打印一行日志。
//type interceptor func(ctx context.Context, handler func(ctx context.Context))

// 将 handler 单独定义成一种类型
type handler func(ctx context.Context)
type interceptor func(ctx context.Context, h handler)

func testInterceptor() {

	var ctx context.Context
	var ceps []interceptor

	var h = func(ctx context.Context) {
		fmt.Println("do something")
	}

	var inter1 = func(ctx context.Context, h handler) {
		fmt.Println("interceptor1")
		h(ctx)
	}
	var inter2 = func(ctx context.Context, h handler) {
		fmt.Println("interceptor2")
		h(ctx)
	}

	ceps = append(ceps, inter1, inter2)
	for _, cep := range ceps {
		cep(ctx, h)
	}

}

type invoker func(ctx context.Context, interceptor2 []interceptor2, h handler) error
type interceptor2 func(ctx context.Context, h handler, ivk invoker) error

func getInvoker(interceptors []interceptor2, cur int, ivk invoker) invoker {
	if cur == len(interceptors)-1 {
		return ivk
	}

	return func(ctx context.Context, interceptors []interceptor2, h handler) error {
		return interceptors[cur+1](ctx, h, getInvoker(interceptors, cur+1, ivk))
	}
}

func getChainInterceptor(interceptors []interceptor2) interceptor2 {
	if len(interceptors) == 0 {
		return nil
	}
	if len(interceptors) == 1 {
		return interceptors[0]
	}
	return func(ctx context.Context, h handler, ivk invoker) error {
		return interceptors[0](ctx, h, getInvoker(interceptors, 0, ivk))
	}

}

func testInterceptor2() {
	var ctx context.Context
	var ceps []interceptor2
	var h = func(ctx context.Context) {
		fmt.Println("do something")
	}

	var inter1 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}
	var inter2 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}
	var inter3 = func(ctx context.Context, h handler, ivk invoker) error {
		h(ctx)
		return ivk(ctx, ceps, h)
	}

	ceps = append(ceps, inter1, inter2, inter3)
	var ivk = func(ctx context.Context, interceptors []interceptor2, h handler) error {
		fmt.Println("invoker start")
		return nil
	}

	cep := getChainInterceptor(ceps) // 将多个handler 合并成一个
	cep(ctx, h, ivk)
}

func main() {
	//testInterceptor()
	//testInterceptor2()
	testMyInterceptor()
}
