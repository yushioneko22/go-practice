package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

var events = []Event{}

func main() {
	e := echo.New()

	// GET: 一覧
	e.GET("/events", func(c echo.Context) error {
		return c.JSON(http.StatusOK, events)
	})

	// POST: 登録
	e.POST("/events", func(c echo.Context) error {
		var ev Event
		if err := c.Bind(&ev); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		ev.ID = len(events) + 1
		events = append(events, ev)
		return c.JSON(http.StatusCreated, ev)
	})

	e.Start(":8080")
}
