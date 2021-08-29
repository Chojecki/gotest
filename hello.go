package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")

	api.GET("/ktojestpanem", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "widzew"})
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()
}
