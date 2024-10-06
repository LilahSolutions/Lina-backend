package main

import (
	"log"
	"os"

	"agrobal-backend/internal/server"
)

// @title crv-middlend
// @version 1.0.0
// @BasePath /api

func main() {
	// When using local development uncomment this line of code with your own port
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	serv.Start()
}
