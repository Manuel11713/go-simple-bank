package main

import (
	"database/sql"
	"log"

	"github.com/Manuel11713/simple-bank/api"
	db "github.com/Manuel11713/simple-bank/db/sqlc"
	"github.com/Manuel11713/simple-bank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot conect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	log.Println("server is running at: ", config.ServerAddress)
}
