package main

import (
	"github.com/andariel0905/expenses-tracker/db"
)

func main() {
	client, cxt, cancel := db.SetupMongoDB()
	defer db.CloseConnection(client, cxt, cancel)
}
