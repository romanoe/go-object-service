package objects

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type ObjectServer struct {
	PgConn *pgxpool.Pool
}

func (t ObjectServer) AddNewObject(ctx echo.Context) error {
	o := &Object{}
	if err := ctx.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdId, _ := CreateObject(t.PgConn, o)
	return ctx.JSON(http.StatusCreated, createdId)
}

func (t ObjectServer) FindObjectByID(ctx echo.Context, id int64) error {
	object, _ := GetObjectById(t.PgConn, id)
	return ctx.JSON(http.StatusOK, object)
}

func (t ObjectServer) DeleteObjectByID(ctx echo.Context, id int64) error {
	deletedId, _ := DeleteObjectById(t.PgConn, id)
	return ctx.JSON(http.StatusOK, deletedId)
}

func (t ObjectServer) FindObjects(ctx echo.Context) error {
	objects, _ := GetAllObjects(t.PgConn)
	return ctx.JSON(http.StatusOK, objects)
}

func NewServer(pool *pgxpool.Pool) *echo.Echo {
	// New router
	e := echo.New()

	// Allow CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Initialize ObjectServer
	objectServer := ObjectServer{PgConn: pool}

	// Register handlers (Router, Server interface)
	RegisterHandlers(e, objectServer)
	e.Logger.Fatal(e.Start(":1323"))

	return e
}
