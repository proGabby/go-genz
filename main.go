package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/presenter/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		log.Println("PORT variable not set")
	}
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	r := mux.NewRouter()

	// Initialize the DB
	db, err := postgressDatasource.InitDatabase()

	if err != nil {
		log.Fatal(err)
	}

	routes.SetUpUserRoutes(r, db)
	log.Fatal(http.ListenAndServe(":8080", r))
}
