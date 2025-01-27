package main

import (
	"database/sql"
	"log"

	"github.com/sgrumley/deskday/internal/service/display"
	"github.com/sgrumley/deskday/internal/store"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "~/.local/share/deskday/network_connections.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ds := store.New(db)
	service := display.New(ds)
	if err := service.Out(); err != nil {
		log.Fatal(err)
	}
}
