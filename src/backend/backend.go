package backend

import (
  "log"
  "github.com/gotk3/gotk3/gtk"
)

var BACKGROUND = [3]int{61, 61, 61}

type button struct {
  w      *gtk.Window
  b      *gtk.Button
  x      int
  y      int
  width  int
  height int
}

type container struct {
  c *gtk.Fixed
}

func (b *button) IsPressed(win *gtk.Window, function func()) {
  b.b.Connect("clicked", func() {
    function()
  })
}

func (b *button) AddContainer(c *container) {
  c.c.Put(b.b, b.x, b.y)
}

func NewButton(win *gtk.Window, label string, x, y, width, height int) button {

  b, err := gtk.ButtonNewWithLabel(label)

  if err != nil {
    log.Fatal("Unable to create a button:", err)
    panic(err)
  }

  // Set button position and size (x, y, width, height) within the fixed container
  b.SetSizeRequest(width, height) // width and height

  return button{win, b, x, y, width, height}
}

func NewContainer(win *gtk.Window) container {
  fixed, err := gtk.FixedNew()
  if err != nil {
    log.Fatal("Unable to create Fixed container:", err)
  }
  win.Add(fixed)

  return container{fixed}
}
