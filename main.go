package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New("stopwatch")
	fmt.Println("kocham rodzine!")
	w := a.NewWindow("Stop Watch")
	w.SetContent(widget.NewLabel("welcome to this horrible app"))
}
