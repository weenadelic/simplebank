package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/weenadelic/simplebank/api"
	db "github.com/weenadelic/simplebank/db/sqlc"
	"github.com/weenadelic/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to postgres")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
