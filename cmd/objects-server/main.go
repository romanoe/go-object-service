package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"object-service/internal/objects"
)

func main() {

	e := echo.New()
	e.GET("/objects", func(ctx echo.Context) error {
		objects, _ := objects.GetAllObjects()
		return ctx.JSON(http.StatusOK, objects)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
