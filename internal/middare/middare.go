// https://github.com/FeifeiyuM/go-microservices-boilerplate/wiki/%E4%B8%89%E3%80%81%E6%9C%8D%E5%8A%A1%E6%A1%86%E6%9E%B6%E4%B8%AD%E9%97%B4%E4%BB%B6%EF%BC%88%E6%8B%A6%E6%88%AA%E5%99%A8%EF%BC%89%E7%9A%84%E5%AE%9E%E7%8E%B0

package main

import (
	"fmt"
)

// 一、中间件实现原理 中间件函数本质上就是一个函数的套娃结构，中间件函数一般都是闭包函数，其接受 handler 作为参数，返回值也是 handler 函数。
// 如下所示
type handlerFunc func(param interface{}) error

// 间件函数外层是一个高阶函数，其接受 handler 函数作为参数，又返回了一个新的 handler 函数。
//中间件内层就是实现了返回了新的 handler 函数，其本身又是一个闭包函数。
func middlewareFunc1(handler handlerFunc) handlerFunc {
	// 返回了一个闭合函数 可以延迟执行
	return func(param interface{}) error {
		// do something func1 before
		fmt.Println("middlewareFunc1 start...")
		err := handler(param)
		// do something func1 after
		fmt.Println("middlewareFunc1 end...")
		return err
	}
}

func middlewareFunc2(handler handlerFunc) handlerFunc {
	// 返回了一个闭合函数 可以延迟执行
	return func(param interface{}) error {
		// do something func1 before
		fmt.Println("middlewareFunc2 start...")
		err := handler(param)
		// do something func1 after
		fmt.Println("middlewareFunc2 end...")
		return err
	}
}

func middlewareFunc3(handler handlerFunc) handlerFunc {
	// 返回了一个闭合函数 可以延迟执行
	return func(param interface{}) error {
		// do something func1 before
		fmt.Println("middlewareFunc3 start...")
		err := handler(param)
		// do something func1 after
		fmt.Println("middlewareFunc3 end...")
		return err
	}
}

func handler(param interface{}) error {
	fmt.Printf("this is the true handler: param %v \n", param)
	return nil
}

//请求的执行顺序一般是 func1 before -> func2 before -> func3 before -> handler(业务逻辑) —> func3 after -> func2 after-> func1 after

func main() {
	// 将真正的 handler 函数用中间件进行封装，以实现相关功能 合并成了一个handler
	// 只是这样的写法有点丑陋
	newHandler := middlewareFunc1(middlewareFunc2(middlewareFunc3(handler)))
	_ = newHandler("test")
}

// gRPC 实现中间件
// grpc 中间件函数的定义为 func() grpc.UnaryServerInterceptor 和上面不太一致 没有入参 不需要传入 grpc.UnaryServerInterceptor 作为参数。
//其主要门道，可以在中间件注册函数 grpc_middleware.ChainUnaryServer() 中发现
//

/*func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 转换层标准的中间件函数结构
		// chainer 函数将 grpc.UnaryServerInterceptor 转换成了我们之前讨论的中间件结构 grpc.UnaryHandler 即是函数的入参格式，也是函数返回格式
		chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}
		// 将中间件函数串连起来
    	// 这边为什么采用数组倒序？
		// 一般情况下，我们期望最先加入中间件数组的函数时最先被调用的（位于洋葱模型的最外层）
		// 函数中传入 handler 函数是最原始的 handler 函数，没有被中间件函数包裹的
 		// 接受最原始的 handler 函数的中间件函数肯定是最后加入中间数组的入参，依次类推，因此中间件数组迭代的时候是倒序的

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, req)
	}
}*/
