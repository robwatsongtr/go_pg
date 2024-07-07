package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/robwatsongtr/go_pg.git/models"
	"github.com/robwatsongtr/go_pg.git/utils"
)

/*
Takes in a database connection and returns the handler. Functional, higher order handling.
Both http.ResponseWriter and *http.Request are required for HTTP handlers
even through the request is not being used. like req res in express I suppose.
*/

// reads in the whole client table from postgres and passes it along to the server as json
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

// takes input of JSON from request body
func CreateClientHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode incoming JSON request body into a Client struct
		var newClient models.Client
		if err := json.NewDecoder(r.Body).Decode(&newClient); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		// insert Client struct into database
		if err := models.CreateClient(db, &newClient); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		acknowledgement := map[string]string{"message": "Client Created Successfully"}
		utils.RespondWithJSON(w, http.StatusOK, acknowledgement)
	}
}

// takes input of JSON from request body
func UpdateClientHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatedClient models.Client
		if err := json.NewDecoder(r.Body).Decode(&updatedClient); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := models.UpdateClient(db, &updatedClient); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		acknowledgement := map[string]string{"message": "Client Updated Successfully"}
		utils.RespondWithJSON(w, http.StatusOK, acknowledgement)
	}
}

// takes the database id from the URL
func DeleteClientHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract client_id from URL path and convert it to an integer
		pathParts := strings.Split(r.URL.Path, "/")
		clientIDStr := pathParts[2]
		clientID, err := strconv.Atoi(clientIDStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid client_id")
		}

		clientToDelete := models.Client{Client_id: clientID}

		if err := models.DeleteClient(db, &clientToDelete); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

		acknowledgement := map[string]string{"message": "Client Deleted Successfully"}
		utils.RespondWithJSON(w, http.StatusOK, acknowledgement)
	}
}
