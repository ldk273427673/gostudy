package main

import (
	"web/controllers"

	"github.com/labstack/echo"
)

func main() {
	// 实例化一个echo对象
	e := echo.New()
	// GET方法
	// 1，接收者是一个*echo.Echo类型，上面实例化的e，就是这一个类型，所以e可以调用这个方法
	// 2，func (*echo.Echo).GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	// 3，GET方法，参数是一个path路径，一个控制器函数HandlerFunc
	// 4，这个控制器函数的定义是HandlerFunc func(Context) error，接受一个context类型的方法
	// 5，这个控制函数，调用c这个接口变量的各类值，如string等
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// 启动服务，开启端口
	// e.Logger.Fatal(e.Start(":80"))

	// 优化版本
	// 访问/，getRoot函数响应结果，访问url http://127.0.0.1/
	e.GET("/", controllers.GetRoot)
	// 访问动态地址，getId函数响应结果，访问url http://127.0.0.1/任意内容
	// e.GET(":id", getId)
	// // 请求多个参数，访问url http://127.0.0.1/user?name=ldk&age=30
	// e.GET("user", getUser)
	// // 静态页面
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":80"))

}

// func getId(c echo.Context) error {
// 	return c.String(200, c.Param("id"))
// }

// func getUser(c echo.Context) error {
// 	return c.String(200, c.QueryParam("name")+" is "+c.QueryParam("age"))
// }
