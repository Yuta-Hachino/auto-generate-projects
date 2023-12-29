package main

import (
	"github.com/Yuta-Hachino/auto-generate-projects/controller/generate/project"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())  // ログ出力
	e.Use(middleware.Recover()) // 復帰用

	e.GET("/", project.Get)
	e.POST("/generate/project", project.Post)
	e.POST("error", func(c echo.Context) error {
		panic("Panic!") //確定でpanicさせてみる
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
