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
	getAllObjects    = "SELECT id, fk_type , created_at FROM objects.object;"
	getObjectById    = "SELECT  id, fk_type , created_at FROM objects.object1;"
	createObject     = "INSERT INTO objects.object (fk_type, created_at) VALUES ($1, $2) RETURNING id;"
	deleteById       = "DELETE FROM objects.object WHERE id=$1 RETURNING id;"
	getObjectTypes   = "SELECT id, value, is_active from objects.object_type;"
	createObjectType = "INSERT INTO objects.object_type (value, is_active) VALUES ($1, $2) RETURNING id;"
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

	// Handle pgx error
	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	fmt.Println("Setting connection to Postgresql")
	return connPool, nil
}

// curl -H "Content-Type: application/json" http://localhost:1323/objects
func GetAllObjects(conn *pgxpool.Pool) ([]*Object, error) {
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

		err := rows.Scan(&object.Id, &object.FkType, &object.CreatedAt)
		if err != nil {
			panic(err)
		}
		objects = append(objects, object)
	}
	return objects, err
}

// curl -H "Content-Type: application/json" http://localhost:1323/objects/1
func GetObjectById(conn *pgxpool.Pool, id int64) (*Object, error) {
	// Initialize object
	var object = new(Object)

	// Executing query
	fmt.Printf("Executing query %s \n", getObjectById)
	err := conn.QueryRow(context.Background(), getObjectById, id).Scan(&object.Id, &object.FkType, &object.CreatedAt)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}
	return object, err
}

// curl -v -XDELETE -H "Content-Type: application/json" 'http://localhost:1323/objects/1'
func DeleteObjectById(conn *pgxpool.Pool, id int64) (int64, error) {
	// Executing query
	fmt.Printf("Executing query %s \n", deleteById)
	var deletedId int64
	err := conn.QueryRow(context.Background(), deleteById, id).Scan(&deletedId)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	return deletedId, nil
}

// curl -XPOST -H "Content-Type: application/json" -d '{"fk_type": 1 }' http://localhost:1323/objects
func CreateObject(conn *pgxpool.Pool, o *Object) (int64, error) {
	// Date time (created_at)
	now := time.Now()

	objectType := o.FkType

	// Create object
	var object = Object{
		Id:        0,
		FkType:    objectType,
		CreatedAt: now,
	}

	// Executing query
	var lastId int64
	fmt.Printf("Executing query %s \n", createObject)
	err := conn.QueryRow(context.Background(), createObject, &object.FkType, &object.CreatedAt).Scan(&lastId)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}
	return lastId, err
}

// curl -H "Content-Type: application/json" http://localhost:1323/objects/types
func GetObjectTypes(conn *pgxpool.Pool) ([]*ObjectType, error) {
	// Initialize object and objects
	var objectTypes []*ObjectType

	// Execute query and add to object
	fmt.Printf("Executing query %s \n", getObjectTypes)
	rows, err := conn.Query(context.Background(), getObjectTypes)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}

	for rows.Next() {
		var objectType = new(ObjectType)

		err := rows.Scan(&objectType.Id, &objectType.Value, &objectType.IsActive)
		if err != nil {
			panic(err)
		}
		objectTypes = append(objectTypes, objectType)
	}
	return objectTypes, err
}

// curl -XPOST -H "Content-Type: application/json" -d '{"value": "mur" }' http://localhost:1323/objects/types
func CreateObjectType(conn *pgxpool.Pool, o *ObjectType) (int64, error) {

	objectValue := o.Value

	// Create object
	var objectType = ObjectType{
		Id:       0,
		Value:    objectValue,
		IsActive: true,
	}

	// Executing query
	var lastId int64
	fmt.Printf("Executing query %s \n", createObjectType)
	err := conn.QueryRow(context.Background(), createObjectType, &objectType.Value, &objectType.IsActive).Scan(&lastId)

	if err != nil {
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}
	return lastId, err
}
