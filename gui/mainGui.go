package gui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"github.com/andariel0905/expenses-tracker/gui/managers"
	"go.mongodb.org/mongo-driver/mongo"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// set general utilities
func CreateQuitButton(myWindow fyne.Window) fyne.Widget {
	button := widget.NewButton("Quit", func() {
		fmt.Println("Closing GUI")
		myWindow.Close()
	})
	return button
}

// Set Toolbar utilities
func createAdminDropdown(client *mongo.Client, cxt context.Context) *widget.Select {
	options := []string{"Expense Categories"}
	return widget.NewSelect(options, func(selected string) {
		switch selected {
		case "Expense Categories":
			managers.ShowExpenseCategoriesWindow(client, cxt)
		}
	})
}

func createInvestmentsButton() *widget.Button {
	return widget.NewButton("Investments", func() {})
}

func createInstallmentsButton() *widget.Button {
	return widget.NewButton("Installments", func() {})
}

// GUI Setup
func StartGUI(client *mongo.Client, cxt context.Context) {
	fmt.Println("Starting GUI")
	myApp := app.New()

	//User Interface declaration
	myApp.Settings().SetTheme(&MyTheme{})

	myWindow := myApp.NewWindow("TrackExp")

	quit := CreateQuitButton(myWindow)

	admin := createAdminDropdown(client, cxt)

	investments := createInvestmentsButton()
	installments := createInstallmentsButton()

	myWindow.SetContent(container.NewBorder(
		container.New(layout.NewHBoxLayout(), admin, investments, installments, layout.NewSpacer(), quit),
		nil,
		nil,
		nil,
		nil,
	))
	myWindow.Resize(fyne.NewSize(600, 700))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}
