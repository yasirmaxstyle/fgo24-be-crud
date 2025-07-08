package main

import (
	"dashboard-backend/routes"
	"dashboard-backend/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//@title RESTful API Contacts
//@version 1.0
//@description RESTful API for contact list
//@BasePath /

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	utils.InitDB()

	routes.CombineRouter(r)

	godotenv.Load()
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	r.Run(port)
	fmt.Printf("program runs on port: %s\n", port)
}
