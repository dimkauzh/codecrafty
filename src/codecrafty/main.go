package main

import (
	"github.com/dimkauzh/codecrafty/src/codecrafty/menu"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Codecrafty")

	w.SetContent(widget.NewLabel("Welcome to Codecrafty!"))
	menu.Menu(w, a)
	w.ShowAndRun()
}
