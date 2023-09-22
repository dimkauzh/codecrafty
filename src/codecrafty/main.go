package main

import (
	bd "github.com/dimkauzh/codecrafty/src/codecrafty/backend"
	menu "github.com/dimkauzh/codecrafty/src/codecrafty/menu"
	fltk "github.com/pwiecz/go-fltk"
)

func main() {
	fltk.InitStyles()
	win := fltk.NewWindow(400, 300)
	win.SetLabel("codecrafty")
	win.SetColor(bd.BACKGROUND)

	codecrafty(win)

	win.End()
	win.Show()
	fltk.Run()
}

func codecrafty(window *fltk.Window) {
	menu.Menu()
}
