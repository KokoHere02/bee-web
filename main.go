package main

import (
	"bee"
	"log"
	"net/http"
)

func main() {
	engine := bee.NEW()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /")
		w.Write([]byte("Hello, BEE!"))
	})
	engine.Run(":8080")

}
