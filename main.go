package main

import (
	"github.com/andariel0905/expenses-tracker/db"
	"github.com/andariel0905/expenses-tracker/global"
	"github.com/andariel0905/expenses-tracker/gui"
)

func main() {
	client, cxt, cancel := db.SetupMongoDB()
	global.SetGlobalVariables(client, cxt)
	gui.StartGUI()
	defer db.CloseConnection(cancel)
}
