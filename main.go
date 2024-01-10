package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/add", Add)
		v1.GET("/update", Add)
		v1.GET("/check", Add)
		v1.GET("/overview", Add)
		v1.DELETE("/delete", Add)

	}

	r.Run()
}

func Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"i": "v2",
	})
}
