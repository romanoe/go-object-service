package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"object-service/internal/objects"
)

type ObjectServer struct {
}

func (t ObjectServer) AddNewObject(ctx echo.Context) error {
	object, _ := objects.GetAllObjects()
	return ctx.JSON(http.StatusOK, object)
}

func (t ObjectServer) FindObjectByID(ctx echo.Context, id int64) error {
	object, _ := objects.GetObjectById(id)
	return ctx.JSON(http.StatusOK, object)
}

func (t ObjectServer) FindObjects(ctx echo.Context) error {
	objects, _ := objects.GetAllObjects()
	return ctx.JSON(http.StatusOK, objects)
}

func NewServer() *echo.Echo {
	// New router
	e := echo.New()
	// Register handlers (Router, Server interface)
	objects.RegisterHandlers(e, ObjectServer{})

	return e
}

func main() {
	server := NewServer()
	server.Logger.Fatal(server.Start(":1323"))
}
