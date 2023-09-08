package main

import (
	"github.com/dimkauzh/codecrafty/src/codecrafty/menu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var a = app.New()
	var w = a.NewWindow("Codecrafty")

	w.Resize(fyne.NewSize(1280, 720))
	w.SetContent(widget.NewLabel("Welcome to Codecrafty!"))

	codecrafty(w, a)

}

func codecrafty(w fyne.Window, a fyne.App) {
	menu.Menu(w, a)
	w.ShowAndRun()
}
