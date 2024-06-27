package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/robwatsongtr/go_pg.git/models"
	"github.com/robwatsongtr/go_pg.git/utils"
)

/*
Takes in a database connection and returns the handler
both http.ResponseWriter and *http.Request are required for HTTP handlers
even through the request is not being used. like req res in express I suppose.
*/
func GetAllClientsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := models.GetClients(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, clients)
	}
}

func CreateClientHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode incoming JSON request body into a Client struct
		var incomingClient models.Client
		if err := json.NewDecoder(r.Body).Decode(&incomingClient); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		// insert Client struct into database
		if err := models.CreateClient(db, &incomingClient); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		acknowledgement := map[string]string{"message": "Client Created Successfully"}
		utils.RespondWithJSON(w, http.StatusOK, acknowledgement)
	}
}
