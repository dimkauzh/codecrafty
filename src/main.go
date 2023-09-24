package main

import (
  "codecrafty/src/menu"
  "fmt"

  "github.com/gotk3/gotk3/gtk"
)

func main() {
  gtk.Init(nil)

  settings, err := gtk.SettingsGetDefault()
  if err != nil {
    fmt.Println("Unable to get default settings:", err)
  }
  settings.SetProperty("gtk-application-prefer-dark-theme", true)

  win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
  if err != nil {
    fmt.Println("Unable to create window:", err)
  }

  win.SetTitle("GTK Example")
  win.Connect("destroy", func() {
    gtk.MainQuit()
  })

  win.SetDefaultSize(800, 600)

  codecrafty(win)


  win.ShowAll()
  gtk.Main()
}

func codecrafty(win *gtk.Window) {
  menu.Menu(win)
}
