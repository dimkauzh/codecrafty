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
	text1 := canvas.NewText("Welcome to Codecrafty!", nil)

	text1.Resize(fyne.NewSize(400, 250))
	text1.Move(fyne.NewPos(20, 20))

	// Project buttons
	button1 := widget.NewButton("Click me", newProject)

	button1.Resize(fyne.NewSize(150, 30))
	button1.Move(fyne.NewPos(100, 100))

	// Menu container
	text := container.NewWithoutLayout(text1, button1)

	// Drawing the menu
	w.SetContent(text)

}

func newProject() {
	fmt.Println("New Project")
}

func openProject() {
	fmt.Println("Open Project")
}
