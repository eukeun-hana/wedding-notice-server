package main

import (

	"log"
	"net/http"
	"os"

	
	"github.com/eukeun-hana/wedding-notice-server/env"
	"github.com/eukeun-hana/wedding-notice-server/httphandler"
	"github.com/eukeun-hana/wedding-notice-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
		
	}

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
	
}
