
package view

import (
	"github.com/rivo/tview"
)

type Assets struct {
	Flex *tview.Flex
	Table *tview.Table
	Log *tview.TextView
//	Status *tview.TextView
}

func NewAssetsLayout() *Assets {
	table := tview.NewTable().
		SetBorders(true)	
	
	log := tview.NewTextView().SetText("ENTER to download Asset").SetTextAlign(tview.AlignCenter)
//	status := tview.NewTextView().SetText("<Status>").SetTextAlign(tview.AlignCenter)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
//			AddItem(status, 0, 2, false).
			AddItem(table, 0, 8, true).
			AddItem(log, 5, 1, false)

	return &Assets{
		Flex: flex,
		Table: table,
		Log: log,
//		Status: status,
	}
}
