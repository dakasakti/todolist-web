package helpers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseJSON struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	fmt.Println(err.Error())
	errorPage := fmt.Sprintf("%d.html", code)
	err = c.File(errorPage)
	if err != nil {
		c.Logger().Error(err)
	}
}
