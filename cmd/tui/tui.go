package main

import (
	// "fmt"
	// "os"

	// "github.com/EngoEngine/ecs"

	// "github.com/therealfakemoot/ecs-rogue/components"
	// "github.com/therealfakemoot/ecs-rogue/systems"

	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewTableRow(t *tview.Table, s []string, idx int) {
	for i, datum := range s {
		t.SetCell(idx, i, tview.NewTableCell(datum))
	}

}
