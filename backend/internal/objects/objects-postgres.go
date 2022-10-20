package objects

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

// SQL queries
const (
	getAllObjects = "SELECT * FROM objects;"
	getObjectById = "SELECT * FROM objects WHERE id=$1;"
	createObject  = "INSERT INTO objects (id, created_at, type) VALUES ($1, $2, $3) RETURNING id;"
	getMaxid      = "SELECT MAX(id) FROM objects;"
	deleteById    = "DELETE FROM objects WHERE id=$1 RETURNING id;"
)

var pgErr *pgconn.PgError

// Database
func SetConnection() (*pgxpool.Pool, error) {
	DbHost := os.Getenv("DBHOST")
	DbPort := os.Getenv("DBPORT")
	DbName := os.Getenv("DBNAME")
	DbUser := os.Getenv("DBUSER")
	DbPassword := os.Getenv("DBPASSWORD")
	DatabaseUrl := "postgres://" + DbUser + ":" + DbPassword + "@" + DbHost + ":" + DbPort + "/" + DbName

	// Open connection
	connPool, err := pgxpool.New(context.Background(), DatabaseUrl)
	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
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
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
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
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}

	// Initialize object
	var object = new(Object)

	// Executing query
	fmt.Printf("Executing query %s \n", getObjectById)
	err = conn.QueryRow(context.Background(), getObjectById, id).Scan(&object.CreatedAt, &object.Type, &object.Id)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}

	return object, err
}

func DeleteObjectById(id int64) (int64, error) {
	// Set connection and (defer) close
	conn, err := SetConnection()
	defer conn.Close()

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	// Executing query
	fmt.Printf("Executing query %s \n", deleteById)
	var deletedId int64
	err = conn.QueryRow(context.Background(), deleteById, id).Scan(&deletedId)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}

	return deletedId, nil
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
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
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
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}

	return err
}
