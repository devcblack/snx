
package view

import (
	"github.com/rivo/tview"
)

type Search struct {
	Flex *tview.Flex
	Form *tview.Form
	Log *tview.TextView
}

func NewSearchLayout() *Search {
	form := tview.NewForm().AddInputField("Keyword", "", 20, nil, nil).
				AddButton("Enter", nil)
	
	log := tview.NewTextView().SetText("Search something with a keyword").SetTextAlign(tview.AlignCenter)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(form, 0, 2, true).
			AddItem(tview.NewBox(), 0, 8, false).
			AddItem(log, 5, 1, false)
	return &Search{
		Flex: flex,
		Form: form,
		Log: log,
	}
}
