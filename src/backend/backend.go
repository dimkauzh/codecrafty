package backend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var BACKGROUND = [3]int{61, 61, 61}

type button struct {
	w      fyne.Window
	b      *widget.Button
	x      int
	y      int
	width  int
	height int
}

func (b *button) IsPressed(function func()) {
	b.b.OnTapped = func() {
		function()
	}
}

func (b *button) Container() *widget.Button {
	return b.b
}

func NewButton(w fyne.Window, label string, x, y, width, height int) button {
	b := widget.NewButton(label, func() {})
	b.Resize(fyne.NewSize(float32(width), float32(height)))
	b.Move(fyne.NewPos(float32(x), float32(y)))

	return button{w, b, x, y, width, height}
}

func DrawContainer(w fyne.Window, c *fyne.Container) {
	w.SetContent(c)
}
