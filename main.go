package main

import (
	"github.com/andariel0905/expenses-tracker/db"
	"github.com/andariel0905/expenses-tracker/gui"
)

func main() {
	client, cxt, cancel := db.SetupMongoDB()
	gui.StartGUI(client, cxt)
	defer db.CloseConnection(client, cxt, cancel)
}
