package objects

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

// DB connection details
const (
	DbHost      = "127.0.0.1"
	DbPort      = "5432"
	DbName      = "objects"
	DbUser      = "postgres"
	DbPassword  = "postgres"
	DatabaseUrl = "postgres://" + DbUser + ":" + DbPassword + "@" + DbHost + ":" + DbPort + "/" + DbName
)

// SQL queries
const (
	getAllObjects = "SELECT * FROM object"
	getObjectById = "SELECT * FROM object WHERE id=$1"
)

// Database
func SetConnection() (*pgxpool.Pool, error) {
	// Open and (defer) close connection
	connPool, err := pgxpool.New(context.Background(), DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Setting connection to Postgresql")
	return connPool, nil
}

func GetAllObjects() ([]*Object, error) {

	// Set connection and (defer) close
	conn, err := SetConnection()
	defer conn.Close()

	// Initialize object and objects
	var objects []*Object

	// Execute query and add to object
	rows, err := conn.Query(context.Background(), getAllObjects)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var object = new(Object)

		err := rows.Scan(&object.CreatedAt, &object.Type, &object.Id)
		if err != nil {
			panic(err)
		}
		objects = append(objects, object)
	}
	return objects, err
}

func GetObjectById(id int64) (*Object, error) {
	// Set connection and (defer) close
	conn, err := SetConnection()
	defer conn.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	// Initialize object
	var object *Object

	// Executing query
	fmt.Printf("Executing query %s", getObjectById)

	err = conn.QueryRow(context.Background(), getObjectById, id).Scan(&object.CreatedAt, &object.Type, &object.Id)

	// Handling error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	return object, err
}
