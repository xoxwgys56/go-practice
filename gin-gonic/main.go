package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func getMessage(c *gin.Context) {
	message := c.Param("msg")
	c.JSON(200, gin.H{
		"message": fmt.Sprint("You sent me to '%s'", message),
	})
}

func populateDatabase(db *gorm.DB) {
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1)
	fmt.Println(product.ID, product.Code)
	db.First(&product, "code = ?", "D41")
	fmt.Println(product)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}
	populateDatabase(db)

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