package main

import "github.com/robwatsongtr/go_pg.git/db"

func main() {
	db.Init()

	select {} // keep program running indefinitely
}
