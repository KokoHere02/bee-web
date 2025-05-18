package main

import (
	"bee"
	"net/http"
)

func main() {
	engine := bee.New()
	engine.GET("/", func(c *bee.Context) {
		c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	engine.Run(":8080")
}
