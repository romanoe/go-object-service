package objects

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
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
	getAllObjects = "SELECT * FROM object;"
	getObjectById = "SELECT * FROM object WHERE id=$1;"
	createObject  = "INSERT INTO object (id, created_at, type) VALUES ($1, $2, $3);"
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
	fmt.Printf("Executing query %s \n", getAllObjects)
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
	var object = new(Object)

	// Executing query
	fmt.Printf("Executing query %s \n", getObjectById)
	err = conn.QueryRow(context.Background(), getObjectById, id).Scan(&object.CreatedAt, &object.Type, &object.Id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return object, err
}

func CreateObject() (*Object, error) {
	// Set connection and (defer) close
	conn, err := SetConnection()
	defer conn.Close()

	// Initialize object
	var object = new(Object)

	// Get max id (id)
	var maxId int64

	// Date time (created_at)
	var now time.Time

	// Type (Arbre, Antenne etc.)
	var objectType string

	// Executing query
	fmt.Printf("Executing query %s \n", createObject)
	err = conn.QueryRow(context.Background(), createObject, maxId, now, objectType).Scan(&object.CreatedAt, &object.Type, &object.Id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return object, err
}
