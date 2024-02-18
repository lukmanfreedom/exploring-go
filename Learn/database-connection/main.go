package main

import "database-connection/config/db"

func main() {
	// set init connection
	DB := db.InitPostgre()

	db.New(DB)
}
