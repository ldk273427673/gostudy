package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", getRoot)
	e.GET("/:id", getId)
	e.GET("/show", getShow)
	e.POST("/struct", getStruct)

	e.Logger.Fatal(e.Start(":80"))
}

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getId(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, id)
}

func getShow(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(200, "team:"+team+",member"+member)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getStruct(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
