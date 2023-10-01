package menu

import (
	"fmt"

	bd "codecrafty/src/backend"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Menu(w fyne.Window) {
	newproject := bd.NewButton(w, "New Project", 50, 50, 200, 50)
	openproject := bd.NewButton(w, "Open Project", 50, 500, 200, 50)

	newproject.IsPressed(newProject)
	openproject.IsPressed(openProject)

	content := container.NewWithoutLayout(newproject.Container(), openproject.Container())

	bd.DrawContainer(w, content)

}

func newProject() {
	fmt.Println("New Project")
}

func openProject() {
	fmt.Println("Open Project")
}
