package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init()  {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "world"})
}

func main() {
	router := gin.New()
	router.GET("/", hello)
	router.Run(":9090")
}
