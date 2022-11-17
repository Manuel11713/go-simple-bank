package main

import (
	"database/sql"
	"log"

	"github.com/Manuel11713/simple-bank/api"
	db "github.com/Manuel11713/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot conect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	log.Println("server is running at: ", serverAddress)
}
