package main

import (
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"fmt"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	ret, err := exec.Command("cmd", "/K", "multipass", "list").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ret))

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
