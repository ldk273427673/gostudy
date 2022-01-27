package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("看源码!")
	fmt.Println("第一步，实例化一个echo类型的对象，使用New方法，会返回一个指针类型")
	e := echo.New()

	fmt.Println(`
	这个对象调用GET方法，GET方法定义为 func (e *Echo) GET(path string, h HandlerFunc, m ...MiddlewareFunc)，可以看到这个方法
	的接收者是一个Echo的指针类型，因为上一步已经实例化了一个echo类型的变量，这个变量就可以调用这个GET方法，然后第一个参数path是路径字符串类型
	第二个参数是一个HandlerFunc的函数类型，第三个参数是可选的参数，在看这个HandlerFunc函数定义， HandlerFunc func(Context) error,这个控制器
	函数是一个接受Context类型参数的，返回error的匿名函数，然后再看这个Context，这个是个接口，	包含类型有String(code int, s string) error
	返回值是一个code和一个字符串，code是200，等同于http.StatusOK,字符串自定义
	`)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(200, id)
	})
	type user struct {
		name string
		id   int
	}
	e.GET("/user", func(c echo.Context) error {
		u1 := user{"登科", 2}
		return c.JSON(200, u1)
	})

	e.GET("1.html", func(c echo.Context) error {
		html := "<html><head><title>tizi365.com</title></head><body>欢迎访问fdfdfdfdfd.com</body></html>"
		return c.HTML(200, html)
	})
	// e.POST("/test", func(c echo.Context) error {
	// 	username := c.FormValue("username")
	// 	html := "调用模板引擎渲染html页面, username"
	// 	return c.HTML(http.StatusOK, html)
	// })

	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":80"))
}
