package main

import (
	"fmt"
	go_middlewares "github.com/owenliang/go-middlewares"
)

// 日志中间件
func Logger(request *go_middlewares.Request) {
	fmt.Println("请求开始")	// before 下一个中间件
	request.Next()
	fmt.Println("请求结束")	// after 下一个中间件
}

// panic捕获中间件
func Recovery(request *go_middlewares.Request) {
	defer func() {
		recover()
		fmt.Println("我确保panic被捕获")
	}()
	request.Next()
}

func main() {
	// 假设来了一个请求
	request := go_middlewares.NewRequest()

	// 注册了一些中间件
	request.RegisterMiddlewares(Logger, Recovery, func(request *go_middlewares.Request) {
		// 业务处理函数作为中间件的最后一环
		fmt.Println("我是业务逻辑")
	})

	// 然后开始处理请求
	request.Next()
}