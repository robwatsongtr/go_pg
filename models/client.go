package models

import "database/sql"

type Client struct {
	ID          int
	Name        string
	Child_name  string
	Email       string
	Billing_amt float32
}

func GetClients(db *sql.DB) ([]Client, error) {

	clients := []Client{}

	return clients, nil
}
