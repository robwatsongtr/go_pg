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
		if err := rows.Scan(&client.ID, &client.Name, &client.Child_name, &client.Email,
			&client.Billing_amt); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}
