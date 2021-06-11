package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
)

var pool *sql.DB

func main() {
	id := flag.Int64("id", 20, "person ID to find")
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	flag.Parse()
	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	if *id == 0 {
		log.Fatal("missing person ID")
	}
	var err error
	pool, err := sql.Open("mysql", *dsn)

}
