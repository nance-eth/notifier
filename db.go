package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type subscription struct{}

type space struct{}

func initDb() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if db, err = sql.Open("sqlite3", "notifier.db"); err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	log.Println("Connected to database successfully")
}
