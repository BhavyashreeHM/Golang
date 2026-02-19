package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
func main() {
	// fmt.Println("GIN")
	router := gin.Default()

	router.GET("/ping", ping)
	router.Run()
}
