package main

import (
	"database/sql"
	"log"

	"github.com/jaeyoung0509/go-banking/api"
	db "github.com/jaeyoung0509/go-banking/db/sqlc"
	"github.com/jaeyoung0509/go-banking/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
