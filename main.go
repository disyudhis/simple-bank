package main

import (
	"context"
	"log"

	"github.com/disyudhis/simplebank/api"
	db "github.com/disyudhis/simplebank/db/sqlc"
	"github.com/disyudhis/simplebank/util"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err);
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err) 
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}