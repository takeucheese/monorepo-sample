package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()
	r.GET("/ping", func (c *gin.Context) {
		message := "pong"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})
	r.Run(":8080")
}
