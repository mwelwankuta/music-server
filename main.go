package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mwelwankuta/music-server/routes"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")

	if len(PORT) == 0 {
		PORT = "8080"
	}

	// declare router
	router := mux.NewRouter()

	routes.Routes(router)

	log.Printf("server started on port %v...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
