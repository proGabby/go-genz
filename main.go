package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	connStr, ok := os.LookupEnv("PORT")

	if !ok {
		log.Println("PORT variable not set")
	}
	if connStr == "" {
		log.Fatal("PORT environment variable not set")
	}

}
