package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"github.com/andariel0905/expenses-tracker/gui/guiutils"
	"github.com/andariel0905/expenses-tracker/gui/managers"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Set Toolbar utilities
func createAdminDropdown() *widget.Select {
	options := []string{"Expense Categories", "Payment Methods"}
	drowpdown := widget.NewSelect(options, func(selected string) {
		switch selected {
		case "Expense Categories":
			managers.ShowExpenseCategoriesWindow()
		case "Payment Methods":
			managers.ShowPaymentMethodsWindow()
		}
	})

	drowpdown.PlaceHolder = "Admin"
	return drowpdown
}

func createInvestmentsButton() *widget.Button {
	return widget.NewButton("Investments", func() {})
}

func createInstalmentsButton() *widget.Button {
	return widget.NewButton("Instalments", func() {})
}

// GUI Setup
func StartGUI() {
	fmt.Println("Starting GUI")
	myApp := app.New()

	//User Interface declaration
	myApp.Settings().SetTheme(&MyTheme{})

	myWindow := myApp.NewWindow("TrackExp")

	quit := guiutils.CreateQuitButton(myWindow, "Exit", "Closing GUI")

	admin := createAdminDropdown()

	investments := createInvestmentsButton()
	installments := createInstalmentsButton()

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
