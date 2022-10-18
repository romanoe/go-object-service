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
	getAllObjects = "SELECT * FROM objects;"
	getObjectById = "SELECT * FROM objects WHERE id=$1;"
	createObject  = "INSERT INTO objects (id, created_at, type) VALUES ($1, $2, $3) RETURNING id;"
	getMaxid      = "SELECT MAX(id) FROM objects;"
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

func CreateObject(o *NewObject) error {
	// Set connection and (defer) close
	conn, err := SetConnection()
	defer conn.Close()

	// Get max id (id)
	var lastId int64
	err = conn.QueryRow(context.Background(), getMaxid).Scan(&lastId)
	maxId := lastId + 1

	if err != nil {
		fmt.Fprintf(os.Stderr, "Getting max id failed: %v\n", err)
		os.Exit(1)
	}

	// Date time (created_at)
	now := time.Now()

	// Type (Arbre, Antenne etc.)
	if o.Type == "" {
		o.Type = "Undefined"
	}
	objectType := o.Type

	// Create object
	var object = Object{
		CreatedAt: now,
		Id:        maxId,
		Type:      objectType,
	}

	// Executing query
	fmt.Printf("Executing query %s \n", createObject)
	err = conn.QueryRow(context.Background(), createObject, &object.Id, &object.CreatedAt, &object.Type).Scan(&lastId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Inserting new object with id %d failed: %v\n", lastId, err)
		os.Exit(1)
	}

	return err
}
