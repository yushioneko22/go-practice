package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/search", func(c echo.Context) error {
		q := c.QueryParam("q")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"query":   q,
			"results": 0,
		})
	})

	e.Start(":1323")
}
