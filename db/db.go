package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm

	"github.com/asuforce/gin-gorm-tutorial/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}

	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
