package handlers

import (
	"net/http"

	"github.com/robwatsongtr/go_pg.git/db"
	"github.com/robwatsongtr/go_pg.git/models"
	"github.com/robwatsongtr/go_pg.git/utils"
)

// both http.ResponseWriter and *http.Request are required for HTTP handlers
// even through the request is not being used. like req res in express I suppose.
func GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := models.GetClients(db.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, clients)
}
