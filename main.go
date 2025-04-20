package main

import (
	"database/sql"
	"log"
  "os"

	_ "github.com/lib/pq"
  "github.com/joho/godotenv"
	"github.com/Nandgopal-R/todoPilot/db/sqlc"
  "github.com/Nandgopal-R/todoPilot/middleware"
  "github.com/Nandgopal-R/todoPilot/router")

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	q := db.New(conn)
	h := middleware.NewHandler(q)

	r := router.Router(h)
	r.Run(":8080")}

