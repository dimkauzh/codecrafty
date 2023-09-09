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

	//text1.Resize(fyne.NewSize(700, 450))
	startText.Move(fyne.NewPos(0, 0))
	startText.TextSize = 85

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
