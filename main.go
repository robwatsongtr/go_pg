package main

import (
	"log"
	"net/http"

	"github.com/robwatsongtr/go_pg.git/db"
	"github.com/robwatsongtr/go_pg.git/handlers"
	"github.com/robwatsongtr/go_pg.git/utils"
)

func main() {
	// get the db connection
	dbConn, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	// final cleanup
	defer dbConn.Close()

	// pass the db connection to the handler which is wrapped by a logger
	http.HandleFunc("/clients", utils.DisplayLog(handlers.GetClientsHandler(dbConn)))

	log.Println("Server Starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
