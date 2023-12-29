package main

import (
	"github.com/Yuta-Hachino/auto-generate-projects/controller/encrypt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())  // ログ出力
	e.Use(middleware.Recover()) // 復帰用

	e.GET("/", encrypt.Get)
	e.POST("/say", encrypt.Post)
	e.POST("error", func(c echo.Context) error {
		panic("Panic!") //確定でpanicさせてみる
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
