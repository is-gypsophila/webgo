package main

import (
	"net/http"
	"webgo/core"
	_ "webgo/core"
)

func main() {
	engine := core.New()

	engine.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>hello core!!!</h1>")
	})

	engine.GET("/hello", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
	})

	engine.POST("/login", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run(":9999")
}
