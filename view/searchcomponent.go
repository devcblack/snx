package view

import (
	"github.com/rivo/tview"
)

type SearchComponent struct {
	Flex *tview.Flex
	Table *tview.Table
	Log *tview.TextView
//	Status *tview.TextView
}

func NewSearchComponentLayout() *SearchComponent {
	table := tview.NewTable().
		SetBorders(true)
	
	log := tview.NewTextView().SetText("ENTER to show Assets || PRESS 'n' to go next site || PRESS 'N' to go previous site").SetTextAlign(tview.AlignCenter)
//	status := tview.NewTextView().SetText("<Status>").SetTextAlign(tview.AlignCenter)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
//			AddItem(status, 0, 2, false).
			AddItem(table, 0, 8, true).
			AddItem(log, 5, 1, false)
	return &SearchComponent{
		Flex: flex,
		Table: table,
		Log: log,
//		Status: status,
	}
}
