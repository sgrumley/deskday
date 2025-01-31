package main

import (
	"database/sql"
	"log"
	"os/user"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sgrumley/deskday/internal/service/display"
	"github.com/sgrumley/deskday/internal/store"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(usr.HomeDir, ".local/share/deskday/network_connections.db")
	db, err := sql.Open("sqlite3", dbPath)
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
