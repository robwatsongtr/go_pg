package router

import (
	"database/sql"
	"net/http"

	"github.com/robwatsongtr/go_pg.git/handlers"
	"github.com/robwatsongtr/go_pg.git/utils"
)

func SetupRoutes(db *sql.DB) {
	http.HandleFunc("/clients", clientsRouter(db))
	http.HandleFunc("/clients/", clientIDRouter(db))
}

func clientsRouter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			utils.DisplayLog(handlers.GetAllClientsHandler(db))(w, r)
		case http.MethodPost:
			utils.DisplayLog(handlers.CreateClientHandler(db))(w, r)
		case http.MethodPut:
			utils.DisplayLog(handlers.UpdateClientHandler(db))(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func clientIDRouter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			utils.DisplayLog(handlers.DeleteClientHandler(db))(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
