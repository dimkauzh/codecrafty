package menu

import (
	"fmt"

	bd "codecrafty/src/backend"

	"github.com/gotk3/gotk3/gtk"
)

func Menu(win *gtk.Window) {

	newproject := bd.NewButton(win, "New Project", 50, 50, 200, 50)

	_ = newproject

}

func newProject() {
	fmt.Println("New Project")
}

func openProject() {
	fmt.Println("Open Project")
}
