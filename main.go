package main

import (
	"fmt"

	mongodb "social-network/infra/provider"
	server "social-network/infra/server"
	getEnv "social-network/infra/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error to load env file")
	}

	port := getEnv.Execute("PORT", "3000")
	engine := gin.Default()

	db, err := mongodb.Connect()
	if err != nil {
		fmt.Println("Database not connected")
	} else {
		fmt.Println("Database connected successfully")
		db.Close()
	}

	server.SetupRoutes(engine)
	engine.Run(":" + port)
}
