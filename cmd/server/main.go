package main

import (
	db "backend/pkg/db"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"backend/internal/healthcheck"
	"backend/internal/sorter"

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

	err = db.Migrate(conn)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Listening on port", os.Getenv("PORT"))
	router := mux.NewRouter()
	router.HandleFunc("/health-check", healthcheck.Handler)
	router.HandleFunc("/classify/{userId}", sorter.ClassifyHandler)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)

}
