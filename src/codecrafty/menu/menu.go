package menu

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Menu(w fyne.Window, a fyne.App) {
	w.SetContent(widget.NewLabel("Test"))

	// Main text
	startText := canvas.NewText("Codecrafty", nil)

	welcome := canvas.NewText("Welcome to Codecrafty!", nil)

	// Setting up the start text
	startText.Move(fyne.NewPos(400, 10))
	startText.TextSize = 85

	// Setting up the welcome text
	welcome.Move(fyne.NewPos(470, 120))
	welcome.TextSize = 25

	// Project buttons
	button1 := widget.NewButton("New Project", newProject)
	button2 := widget.NewButton("Open Project", openProject)

	// Setting up the new project button
	button1.Resize(fyne.NewSize(200, 50))
	button1.Move(fyne.NewPos(150, 250))

	// Setting up the open project button
	button2.Resize(fyne.NewSize(200, 50))
	button2.Move(fyne.NewPos(150, 320))

	// Drawing with a menu container
	w.SetContent(container.NewVBox(startText, welcome, button1, button2))

}

func newProject() {
	fmt.Println("New Project")
}

func openProject() {
	fmt.Println("Open Project")
}
