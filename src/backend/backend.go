package backend

import (
  "fmt"

  "github.com/gotk3/gotk3/gtk"
)

var BACKGROUND = [3]int{61, 61, 61}

type button struct {
  w      *gtk.Window
  b      *gtk.Button
  c      *gtk.Container
  x      int
  y      int
  width  int
  height int
}

func (b *button) IsPressed(win *gtk.Window, function func()) {
  b.b.Connect("clicked", func() {
    function()
  })
}

func NewButton(win *gtk.Window, label string, x, y, width, height int) button {

  // Create a fixed container
  fixed, err := gtk.FixedNew()
  if err != nil {
    fmt.Println("Unable to create Fixed container:", err)
  }
  win.Add(fixed)

  b, err := gtk.ButtonNewWithLabel(label)

  if err != nil {
    fmt.Println("Unable to create a button:", err)
  }

  // Set button position and size (x, y, width, height) within the fixed container
  b.SetSizeRequest(width, height) // width and height
  fixed.Put(b, x, y)              // x and y

  return button{win, b, nil, x, y, width, height}
}
