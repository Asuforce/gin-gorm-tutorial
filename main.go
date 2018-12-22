package main

import (
	"github.com/gin-gonic/gin"

	"github.com/asuforce/gin-gorm-tutorial/db"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	db.Init()
	r.Run()
}
