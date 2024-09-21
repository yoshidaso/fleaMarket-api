package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//items := []models.Item{
	//	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", soldOut: false}
	//}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
