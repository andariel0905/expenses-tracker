package guiutils

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateQuitButton(myWindow fyne.Window, printlnText string) fyne.Widget {
	button := widget.NewButton("Quit", func() {
		fmt.Println(printlnText)
		myWindow.Close()
	})
	return button
}
