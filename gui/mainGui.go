package gui

// Code from https://betterprogramming.pub/how-to-create-a-simple-data-entry-desktop-app-with-golang-and-fyne-7c9e543d71e

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"github.com/andariel0905/expenses-tracker/handlers"
	"go.mongodb.org/mongo-driver/mongo"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/data/binding"
)

func StartGUI(client *mongo.Client, cxt context.Context) {
	fmt.Println("Starting GUI")
	myApp := app.New()
	myWindow := myApp.NewWindow("TrackExp")

	loadedExpenseCategories := handlers.GetExpenseCategories(client, cxt)
	var expenseCategoriesNames []string

	for _, document := range loadedExpenseCategories {
		name, ok := document["Name"].(string)

		if !ok {
			println("No property 'Name' in the document or it is not a string")
			continue
		}
		expenseCategoriesNames = append(expenseCategoriesNames, name)
	}

	data := binding.NewStringList()
	data.Set(expenseCategoriesNames)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	list.OnSelected = func(id widget.ListItemID) {
		list.Unselect(id)
		d, _ := data.GetValue(id)
		w := myApp.NewWindow("Edit Data")

		itemName := widget.NewEntry()
		itemName.Text = d

		updateData := widget.NewButton("Update", func() {
			data.SetValue(id, itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		deleteData := widget.NewButton("Delete", func() {
			var newData []string
			dt, _ := data.Get()

			for index, item := range dt {
				if index != id {
					newData = append(newData, item)
				}
			}

			data.Set(newData)

			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, updateData, deleteData, cancel))
		w.Resize(fyne.NewSize(400, 200))
		w.CenterOnScreen()
		w.Show()

	}

	add := widget.NewButton("Add", func() {
		w := myApp.NewWindow("Add Data")

		itemName := widget.NewEntry()

		addData := widget.NewButton("Add", func() {
			data.Append(itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, addData, cancel))
		w.Resize(fyne.NewSize(400, 200))
		w.CenterOnScreen()
		w.Show()

	})

	exit := widget.NewButton("Quit", func() {
		fmt.Println("Closing GUI")
		myWindow.Close()
	})

	myWindow.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}
