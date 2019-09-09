package main

import (
	"github.com/asuforce/gin-gorm-tutorial/db"
	"github.com/asuforce/gin-gorm-tutorial/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
