package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. Echoインスタンスの生成
	e := echo.New()

	// 2. ミドルウェアの導入 (ログの表示)
	e.Use(middleware.Logger())

	// 3. ルーティング
	e.GET("/data", func(c echo.Context) error {
		// クエリパラメータの取得
		q := c.QueryParam("q")

		// JSONレスポンス
		return c.JSON(http.StatusOK, map[string]interface{}{
			"query": q,
			"type":  "search",
		})
	})

	// 4. サーバー起動 (ポート1323がデフォルト)
	e.Logger.Fatal(e.Start(":1323"))
}
