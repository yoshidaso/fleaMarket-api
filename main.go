package main

import (
	"gin-freemarket/models"
	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", soldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", soldOut: true},
		{ID: 1, Name: "商品3", Price: 3000, Description: "説明3", soldOut: false},
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
