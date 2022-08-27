package main

import (
	db "backend/pkg/db"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"backend/internal/healthcheck"

	"github.com/gorilla/mux"
)

var flagConfig = flag.String("config", "./config/local.env", "path to config file")

func main() {
	log.Println("Loading config vars")
	flag.Parse()

	err := godotenv.Load(*flagConfig)

	if err != nil {
		log.Fatal("Failed to load config file")
	}

	conn, err := db.Connection()

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	db.Migrate(conn)

	log.Println("Starting api server with", os.Getenv("MESSAGE"))
	router := mux.NewRouter()
	router.HandleFunc("/health-check", healthcheck.Handler)

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)

}
