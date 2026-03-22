package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. デフォルトの設定でエンジンを作成
	r := gin.Default()

	// 2. GETメソッドの定義
	r.GET("/user/:id", func(c *gin.Context) {
		// パスパラメータの取得
		id := c.Param("id")

		// JSONで返却 (状態コード200)
		c.JSON(200, gin.H{
			"user_id": id,
			"status":  "active",
		})
	})

	// 3. サーバー起動 (デフォルトは :8080)
	r.Run()
}
