package main

import (
	config "GOLANG_CRUD/Config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
