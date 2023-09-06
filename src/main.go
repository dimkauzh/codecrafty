package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Codecrafty")

	w.SetContent(widget.NewLabel("Welcome to Codecrafty!"))
	w.ShowAndRun()
}
