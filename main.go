package main

import (
	"duomly.com/go-bank-backend/api"
	"duomly.com/go-bank-backend/database"
	"duomly.com/go-bank-backend/migrations"
)

func main() {
	// Do migration
	// migrations.MigrateTransactions()

	// Init database
	database.InitDatabase()
	migrations.Migrate()
	api.StartApi()

}
