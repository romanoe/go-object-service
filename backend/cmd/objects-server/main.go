package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"object-service/internal/objects"
)

type ObjectServer struct {
}

func (t ObjectServer) AddNewObject(ctx echo.Context) error {
	o := &objects.NewObject{}
	if err := ctx.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := objects.CreateObject(o)
	return ctx.JSON(http.StatusCreated, err)
}

func (t ObjectServer) FindObjectByID(ctx echo.Context, id int64) error {
	object, _ := objects.GetObjectById(id)
	return ctx.JSON(http.StatusOK, object)
}

func (t ObjectServer) DeleteObjectByID(ctx echo.Context, id int64) error {
	deletedId, _ := objects.DeleteObjectById(id)
	return ctx.JSON(http.StatusOK, deletedId)
}

func (t ObjectServer) FindObjects(ctx echo.Context) error {
	objects, _ := objects.GetAllObjects()
	return ctx.JSON(http.StatusOK, objects)
}

func NewServer() *echo.Echo {
	// New router
	e := echo.New()

	// Allow CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Register handlers (Router, Server interface)
	objects.RegisterHandlers(e, ObjectServer{})

	return e
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := NewServer()
	server.Logger.Fatal(server.Start(":1323"))
}
