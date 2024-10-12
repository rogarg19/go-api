package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitgarg19/go-api/models"
)

func main() {
	r := gin.Default()

	r.GET("/", HandleRoot)
	r.GET("/albums", GetAlbums)

	r.Run()
}

func HandleRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"hello": "there",
	})
}

func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, models.Albums)
}
