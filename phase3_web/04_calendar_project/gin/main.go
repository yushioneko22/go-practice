package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Date  string `json:"date" binding:"required"`
}

var events = []Event{}

func main() {
	r := gin.Default()

	// GET: 一覧
	r.GET("/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, events)
	})

	// POST: 登録
	r.POST("/events", func(c *gin.Context) {
		var e Event
		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		e.ID = len(events) + 1
		events = append(events, e)
		c.JSON(http.StatusCreated, e)
	})

	r.Run(":8080")
}
