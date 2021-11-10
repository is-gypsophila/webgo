package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 定义map的别名为H
type H map[string]interface{}

// Context 定义上下文结构体
type Context struct {
	// 返回对象
	Writer http.ResponseWriter
	// 请求对象
	Req *http.Request
	// 访问路径
	Path string
	// 访问方法
	Method string
	// 状态码
	StatusCode int
}

// 创建Context对象
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
