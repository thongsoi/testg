package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/testg/database"
	"github.com/thongsoi/testg/internal/order"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Println("Failed to close the database:", err)
		}
	}()

	r := mux.NewRouter()

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", order.OrderHandler).Methods("GET")

	// HTMX route for fetching market options
	r.HandleFunc("/markets", order.FetchMarketsHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
