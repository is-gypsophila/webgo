package main

import (
	"net/http"
	"webgo/core"
	_ "webgo/core"
)

func main() {
	r := core.New()
	r.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *core.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *core.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
