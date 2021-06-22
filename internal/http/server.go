package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
