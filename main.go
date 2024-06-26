package main

import (
	"log"
	"net/http"

	"github.com/robwatsongtr/go_pg.git/db"
	"github.com/robwatsongtr/go_pg.git/handlers"
	"github.com/robwatsongtr/go_pg.git/utils"
)

func main() {
	db.Init()
	defer db.DB.Close()

	http.HandleFunc("/clients", utils.DisplayLog(handlers.GetClientsHandler))

	log.Println("Server Starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
