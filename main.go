package main

import (
	"log"

	"github.com/depinsuthap/elabram-backend-test/config"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	config.Route()
}
