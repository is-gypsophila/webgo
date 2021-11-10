package core

import (
	"net/http"
)

// HandlerFunc 定义请求处理函数
type HandlerFunc func(c *Context)

// Engine 引擎结构体
type Engine struct {
	router *router
}

// New 创建引擎实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 给引擎Engine添加addRoute方法
func (engine Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
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

// 自定义ServeHTTP接口实现
func (engine Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.handle(c)
}
