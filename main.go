package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/Nandgopal-R/todoPilot/db/sqlc"
  "github.com/Nandgopal-R/todoPilot/middleware"
  "github.com/Nandgopal-R/todoPilot/router")

func main() {
	conn, err := sql.Open("postgres", "your_connection_string")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	q := db.New(conn)
	h := middleware.NewHandler(q)

	r := router.SetupRouter(h)
	r.Run(":8080")
}

