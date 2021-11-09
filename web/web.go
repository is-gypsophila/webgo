package web

import (
	"fmt"
	"net/http"
)

// HandlerFunc 请求处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 引擎结构体
type Engine struct {
	router map[string]HandlerFunc
}

// New 创建引擎实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 给引擎Engine添加addRoute方法
func (engine Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 定义get函数
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义post函数
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动服务器监听端口
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	}
}
