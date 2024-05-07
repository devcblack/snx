package view

import (
	"github.com/rivo/tview"
)

type Login struct {
	Flex *tview.Flex
	Form *tview.Form
	Log *tview.TextView
}

func NewLoginLayout() *Login {
	form := tview.NewForm().AddInputField("Username", "", 20, nil, nil).
		AddPasswordField("Password", "", 20, 42, nil).
		AddInputField("Url", "", 20, nil, nil).
		AddButton("Enter", nil)
	
	log := tview.NewTextView().SetText("Welcome").SetTextAlign(tview.AlignCenter)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(form, 0, 2, true).
			AddItem(tview.NewBox(), 0, 8, false).
			AddItem(log, 5, 1, false)
	return &Login{
		Flex: flex,
		Form: form,
		Log: log,
	}
}
