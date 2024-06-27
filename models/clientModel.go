package models

import "database/sql"

type Client struct {
	Client_id   int     `json:"client_id"`
	Name        string  `json:"name"`
	Child_name  string  `json:"child_name"`
	Email       string  `json:"email"`
	Billing_amt float32 `json:"billing_amt"`
}

func GetClients(db *sql.DB) ([]Client, error) {
	// read in the rows from the clients table, error handling
	rows, err := db.Query("SELECT client_id, name, child_name, email, billing_amt FROM clients")
	if err != nil {
		return nil, err
	}
	// when function is finished executing perform cleanup and close
	defer rows.Close()

	// initializes an empty 'slice' of Client objects
	clients := []Client{}

	// As long as there's a row, read the row into a client object and save into client slice
	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.Client_id, &client.Name, &client.Child_name, &client.Email,
			&client.Billing_amt); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func CreateClient(db *sql.DB, client *Client) error {
	query := `INSERT INTO clients (name, child_name, email, billing_amt) VALUES ($1, $2, $3, $4) 
		RETURNING client_id`
	err := db.QueryRow(query, client.Name, client.Child_name,
		client.Email, client.Billing_amt).Scan(&client.Client_id)

	return err
}

func UpdateClient(db *sql.DB, client *Client) error {
	query := "UPDATE clients SET name=$1, child_name=$2, email=$3, billing_amt=$4 WHERE client_id=$5"
	_, err := db.Exec(query, client.Name, client.Child_name, client.Email,
		client.Billing_amt, client.Client_id)

	return err
}
