package main

import (
	"fmt"
	"net/http"
	"webgo/web"
	_ "webgo/web"
)

func main() {
	engine := web.New()

	engine.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	engine.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	engine.Run(":9999")
}
