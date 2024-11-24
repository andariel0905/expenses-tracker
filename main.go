package main

import (
	"github.com/andariel0905/expenses-tracker/db"
)

func main() {
	client, context, cancel := db.SetupMongoDB()
	defer db.CloseConnection(client, context, cancel)
}