// Package main provides various examples of Fyne API capabilities.
package main

import (
	_ "GG/Server"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//Server.Getdata()

	a := app.New()
	w := a.NewWindow("Notepad")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Ежедневник")

	name_label := widget.NewLabel("Название дела")
	name := widget.NewEntry()

	case_description := widget.NewLabel("Описание дела")
	description := widget.NewEntry()

	date := widget.NewLabel("Дата")
	date_r := widget.NewEntry()

	reminder := widget.NewLabel("Напомнить?")
	reminder_r := widget.NewRadioGroup([]string{"Да", "Нет"}, func(string) {
		fyne.CurrentApp().SendNotification(&fyne.Notification{Title: "Fyne Demo", Content: "Testing notifications..."})
	})

	result := widget.NewLabel("")

	btn := widget.NewButton("Назначить", func() {
		username := name.Text
		description_r := description.Text
		male_val := date_r.Text
		rem_arr := reminder_r.Selected
		result.SetText(username + "\n" + description_r + "\n" + male_val + "\n" + rem_arr + "\n")

	})

	w.SetContent(container.NewVBox(
		title,
		name_label,
		name,
		case_description,
		description,
		date,
		date_r,
		reminder,
		reminder_r,
		btn,
		result,
	))

	w.ShowAndRun()
}
