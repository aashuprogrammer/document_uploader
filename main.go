package main

import (
	"context"
	"log"

	"github.com/aashuprogrammer/document_uploader.git/api"
	"github.com/aashuprogrammer/document_uploader.git/db/pgdb"
	"github.com/aashuprogrammer/document_uploader.git/token"
	"github.com/aashuprogrammer/document_uploader.git/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// config
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to read env", err)
	}

	// db
	poolConfig, err := pgxpool.ParseConfig(config.DatabaseURL)
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	dbConn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}
	store := pgdb.NewStore(dbConn)
	defer dbConn.Close()

	// token
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("failed to create token maker ", err)
	}

	// fiber server
	server, err := api.NewServer(config, store, tokenMaker)

	if err != nil {
		log.Fatal("failed to create server ", err)
	}
	err = server.Start(config.Port)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
