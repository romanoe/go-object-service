package main

import (
	"github.com/joho/godotenv"
	"log"
	"object-service/internal/objects"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Open (and defer close) postgres connection (defer close())
	conn, err := objects.SetConnection()
	defer conn.Close()

	if err != nil {
		log.Fatal("Error connecting to postgres")
	}

	// Create new server and use postgres connection
	objects.NewServer(conn)

}
