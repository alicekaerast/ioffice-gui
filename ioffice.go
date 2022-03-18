package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/alicekaerast/ioffice/lib"
	"time"
)

func main() {
	a := app.NewWithID("info.kaerast.ioffice")
	w := a.NewWindow("iOffice")

	txtUsername := widget.NewEntry()
	txtUsername.SetText(a.Preferences().String("Username"))
	txtPassword := widget.NewPasswordEntry()
	txtPassword.SetText(a.Preferences().String("Password"))
	txtHostname := widget.NewEntry()
	txtHostname.SetText(a.Preferences().String("Hostname"))

	form := &widget.Form{OnSubmit: func() {
		a.Preferences().SetString("Username", txtUsername.Text)
		a.Preferences().SetString("Password", txtPassword.Text)
		a.Preferences().SetString("Hostname", txtHostname.Text)

		ioffice := lib.NewIOffice(txtHostname.Text, txtUsername.Text, txtPassword.Text, "")

		reservations := ioffice.GetReservations()

		list := widget.NewList(
			func() int {
				return len(reservations) + 1
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				if i == 0 {
					o.(*widget.Label).SetText("Location | Time")
				} else {
					o.(*widget.Label).SetText(fmt.Sprintf("%v | %v", reservations[i-1].Room.Name, time.Unix(reservations[i-1].StartDate/1000, 0).String()))
				}
			})
		w.SetContent(list)
	},
		SubmitText: "Show Reservations",
	}

	form.Append("Username:", txtUsername)
	form.Append("Password:", txtPassword)
	form.Append("Hostname:", txtHostname)

	grid := container.New(layout.NewVBoxLayout(), form)

	w.SetContent(grid)

	w.Resize(fyne.Size{Width: 600, Height: 600})
	w.CenterOnScreen()
	w.ShowAndRun()
}
