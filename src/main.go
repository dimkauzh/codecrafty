package main

import (
	"codecrafty/src/menu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("codecrafty")

	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow.Resize(fyne.NewSize(800, 600))

	codecrafty(myWindow)

	myWindow.ShowAndRun()
}

func codecrafty(w fyne.Window) {
	menu.Menu(w)
}
