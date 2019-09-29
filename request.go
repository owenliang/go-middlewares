package go_middlewares

type HandleFunc func(*Request)

type Request struct {
	index int
	middlewares []HandleFunc	// 中间件 + 处理函数
}

// 生成请求
func NewRequest() (request *Request) {
	request = &Request{
		index:       0,
		middlewares: make([]HandleFunc, 0),
	}
	return
}

// 注册中间件
func (request *Request) RegisterMiddlewares(middlewares ...HandleFunc) {
	for _, mid := range middlewares {
		request.middlewares = append(request.middlewares, mid)
	}
}

// 执行中间件
func (request *Request) Next() {
	index := request.index
	if index >= len(request.middlewares) {
		return
	}

	request.index++
	request.middlewares[index](request)
}