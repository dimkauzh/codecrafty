package menu

import (
	"fmt"

	bd "codecrafty/src/backend"

	"github.com/gotk3/gotk3/gtk"
)

func Menu(win *gtk.Window) {
  container := bd.NewContainer(win)
  
	newproject := bd.NewButton(win, "New Project", 50, 50, 200, 50)
  openproject := bd.NewButton(win, "Open Project", 50, 100, 200, 50)

  newproject.IsPressed(win, newProject)
  openproject.IsPressed(win, openProject)

  newproject.AddContainer(&container)
  openproject.AddContainer(&container)
}

func newProject() {
	fmt.Println("New Project")
}

func openProject() {
	fmt.Println("Open Project")
}
