package encrypt

import (
	"net/http"

	"github.com/Yuta-Hachino/auto-generate-projects/service/chatgpt"

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
	message, err := chatgpt.SendMessage("ESPv2とGoのAPIをCloudRunサービスとしてデプロイしたい。terraformで書くとどうなる？必要なファイルのリストをカンマ区切り文字列にして教えて？その一行以外の説明などはいらない。")

	if err != nil {
		return echo.ErrInternalServerError
	}
	data := ResponseData{
		Status:  200,
		Message: message,
	}
	return c.JSON(http.StatusOK, data)
}

func Post(c echo.Context) error {
	data := new(RequestData)
	if err := c.Bind(&data); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
