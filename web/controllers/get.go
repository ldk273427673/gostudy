package controllers

import "github.com/labstack/echo"

func GetRoot(c echo.Context) error {
	return c.String(200, "this is web root")
}
