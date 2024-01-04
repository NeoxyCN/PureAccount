package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/add", Add)
	}

	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run()
}

func Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"i": "v2",
	})
}
