package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
// curl http://localhost:9999/world
// 404 NOT FOUND: /world

import (
	"net/http"

	"gee"
)

func main() {

	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
