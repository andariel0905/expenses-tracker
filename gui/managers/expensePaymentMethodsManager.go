package managers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"github.com/andariel0905/expenses-tracker/gui/guiutils"
	"github.com/andariel0905/expenses-tracker/handlers"

	"fyne.io/fyne/v2/widget"
)

func getLoadedPaymentMethods() []string {
	loadedPaymentMethods := handlers.GetPaymentMethods()
	var paymentMethodsNames []string

	for _, document := range loadedPaymentMethods {
		name, ok := document["name"].(string)

		if !ok {
			println("No property 'Name' in the document or it is not a string")
			continue
		}
		paymentMethodsNames = append(paymentMethodsNames, name)
	}
	return paymentMethodsNames
}

func loadPaymentMethods() binding.StringList {
	loadedPaymentMethods := getLoadedPaymentMethods()
	data := binding.NewStringList()
	data.Set(loadedPaymentMethods)
	return data
}

func createPaymentMethodsList(myApp fyne.App, data binding.StringList) fyne.Widget {
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

		currentName := d
		itemName := widget.NewEntry()
		itemName.Text = d

		updateData := widget.NewButton("Update", func() {
			handlers.SetPaymentMethod(currentName, itemName.Text)
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

			handlers.DeletePaymentMethod(currentName)
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

func addPaymentMethodWindow(data binding.StringList) {
	w := fyne.CurrentApp().NewWindow("Add Payment Method")

	itemName := widget.NewEntry()

	addData := widget.NewButton("Add", func() {
		handlers.PostPaymentMethod(itemName.Text)
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

func ShowPaymentMethodsWindow() {
	myApp := fyne.CurrentApp()
	myWindow := myApp.NewWindow("Payment Methods Manager")

	data := loadPaymentMethods()
	list := createPaymentMethodsList(myApp, data)

	add := widget.NewButton("New Payment Method", func() {
		addPaymentMethodWindow(data)
	})

	quit := guiutils.CreateQuitButton(myWindow, "Close", "Closing Payment Methods Window")

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
