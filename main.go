package main

import (
	"net/http"
	// 3rd party modules
	"github.com/gin-gonic/gin"
)

//
// Silly functions for managing HTTP methods
//

func SillyGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"object":"Hi! I'm a silly object"})
}

func SillyPost(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message":"Silly object created!"})
}

func SillyPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"Silly object modified!"})
}

func SillyDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"Silly object deleted!"})
}

//
// Main  
//

func main() {
	// setup gin engine
	gin.DisableConsoleColor()
	router := gin.Default()	
	router.GET("/", SillyGet)
	router.POST("/", SillyPost)
	router.PUT("/", SillyPut)
	router.DELETE("/", SillyDelete)

	// listen and serve http requests
	router.Run() 
}