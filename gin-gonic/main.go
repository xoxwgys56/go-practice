package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	message := c.Param("msg")
	c.JSON(200, gin.H{
		"message": fmt.Sprint("You sent me to '%s'", message),
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/message/:msg", getMessage)
	// if call `.Run` without parameter. default port is 8080
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}