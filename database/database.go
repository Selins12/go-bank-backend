package database

import (
	"duomly.com/go-bank-backend/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
func InitDatabase() {
	database, err := gorm.Open("postgres", "postgres://fuwbktzjtanemf:6817f43b5b3b61240b6cf7517a30a3bf76299920b0ff66bd04a3015a31d95db6@ec2-34-239-81-70.compute-1.amazonaws.com:5432/dai8v2i11qupuu?sslmode=require")
	helpers.HandleErr(err)
	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database

}
