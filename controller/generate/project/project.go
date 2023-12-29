package project

import (
	"net/http"

	"github.com/Yuta-Hachino/auto-generate-projects/service/generate"
	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RequestData struct {
	Message string `json:"message"`
}

func Get(c echo.Context) error {
	data := ResponseData{
		Status:  200,
		Message: "message",
	}
	return c.JSON(http.StatusOK, data)
}

func Post(c echo.Context) error {
	data := new(RequestData)
	if err := c.Bind(&data); err != nil {
		return err
	}
	err := generate.NewProject(data.Message)
	if err != nil {
		return echo.ErrInternalServerError
	}
	res := ResponseData{
		Status:  200,
		Message: "url",
	}

	return c.JSON(http.StatusOK, res)
}
