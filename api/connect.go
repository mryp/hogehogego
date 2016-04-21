package api

import (
	"net/http"

	"github.com/labstack/echo"
)

//---------------------------------------------------------------------
//構造体
//---------------------------------------------------------------------
//
type connectResponce struct {
	Result int `json:"result"`
}

//---------------------------------------------------------------------
//関数
//---------------------------------------------------------------------

func ConnectHandler(c echo.Context) error {
	//termid := c.QueryParam("teamid")
	//status := c.QueryParam("status")
	result := connectResponce{Result: -1}

	return c.JSON(http.StatusOK, result)
}
