package main

import (
	"log"
	"net/http"

	"github.com/robwatsongtr/go_pg.git/db"
	"github.com/robwatsongtr/go_pg.git/router"
)

func main() {
	// get the db connection
	db, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	// final cleanup
	defer db.Close()

	// router takes care of different HTTP methods and logging. pass db conn to it.
	router.SetupRoutes(db)

	log.Println("Server Starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
