package main

import (
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
