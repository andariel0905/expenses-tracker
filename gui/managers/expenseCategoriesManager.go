package managers

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/andariel0905/expenses-tracker/gui/guiutils"
	"github.com/andariel0905/expenses-tracker/handlers"
	"go.mongodb.org/mongo-driver/mongo"

	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/data/binding"
)

func getLoadedExpenseCategories(client *mongo.Client, cxt context.Context) []string {
	loadedExpenseCategories := handlers.GetExpenseCategories(client, cxt)
	var expenseCategoriesNames []string

	for _, document := range loadedExpenseCategories {
		name, ok := document["name"].(string)

		if !ok {
			println("No property 'Name' in the document or it is not a string")
			continue
		}
		expenseCategoriesNames = append(expenseCategoriesNames, name)
	}
	return expenseCategoriesNames
}

func loadExpenseCategories(client *mongo.Client, cxt context.Context) binding.StringList {
	loadedExpenseCategories := getLoadedExpenseCategories(client, cxt)
	data := binding.NewStringList()
	data.Set(loadedExpenseCategories)
	return data
}

func createList(myApp fyne.App, data binding.StringList) fyne.Widget {
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
			originalData, _ := data.Get()

			for index, item := range originalData {
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
	return list
}

func addExpenseCategoryWindow(client *mongo.Client, cxt context.Context, data binding.StringList) {
	w := fyne.CurrentApp().NewWindow("Add Expense Category")

	itemName := widget.NewEntry()

	addData := widget.NewButton("Add", func() {
		handlers.PostExpenseCategory(client, cxt, itemName.Text)
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
}

func ShowExpenseCategoriesWindow(client *mongo.Client, cxt context.Context) {
	myApp := fyne.CurrentApp()
	myWindow := myApp.NewWindow("Expense Categories")

	data := loadExpenseCategories(client, cxt)
	list := createList(myApp, data)

	add := widget.NewButton("New Expense Category", func() {
		addExpenseCategoryWindow(client, cxt, data)
	})

	quit := guiutils.CreateQuitButton(myWindow, "Close", "Closing Expense Categories Window")

	myWindow.SetContent(container.NewBorder(
		container.New(layout.NewHBoxLayout(), add, layout.NewSpacer(), quit),
		nil,
		nil,
		nil,
		list,
	))
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.CenterOnScreen()
	myWindow.Show()
}
