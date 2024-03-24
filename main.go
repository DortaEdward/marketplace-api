package main

import (
	"os"

	"github.com/dortaedward/marketplace-api-chi/types"
	"github.com/joho/godotenv"
)

func main() {
	// Loading Env Variables
	err := godotenv.Load()
	if err != nil {
		panic("ERROR: Cannot load env vars!")
	}
	addr := os.Getenv("PORT")

	// Create and Run Server
	server := types.NewServer(addr)
	server.Run()
}
