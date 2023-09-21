package main

import (
	fltk "github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(400, 300)
	win.SetLabel("Main Window")

	codecrafty(win)

	win.End()
	win.Show()
	fltk.Run()
}

func codecrafty(window *fltk.Window) {

}
