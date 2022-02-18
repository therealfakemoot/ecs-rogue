package ui

import (
	"fmt"
	"strings"
	// "os"

	"github.com/therealfakemoot/ecs-rogue/components"

	// "github.com/EngoEngine/ecs"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Record struct {
	Icon         rune
	IconColor    tcell.Color
	Message      string
	MessageColor tcell.Color
}

type ActivityLog struct {
	tview.TableContentReadOnly
	Records []Record
}

func NewActivityLog() *ActivityLog {
	var al ActivityLog
	al.Records = make([]Record, 0)
	return &al
}

// Return the cell at the given position or nil if there is no cell. The
// row and column arguments start at 0 and end at what GetRowCount() and
// GetColumnCount() return, minus 1.
func (al *ActivityLog) GetCell(row, column int) *tview.TableCell {
	r := al.Records[row]
	switch column {
	// TODO: don't forget to apply color here
	case 0:
		return tview.NewTableCell(string(r.Icon))
	case 1:
		return tview.NewTableCell(string(r.Message))
	}
	return tview.NewTableCell("⊕")
}

// Return the total number of rows in the table.
func (al *ActivityLog) GetRowCount() int {
	return len(al.Records)
}

// Return the total number of columns in the table.
func (al *ActivityLog) GetColumnCount() int {
	return 2
}

// Remove all table data.
func (al *ActivityLog) Clear() {}

type Base struct {
	HP          *tview.TextView
	Mana        *tview.TextView
	Weapon      *tview.TextView
	Progression *tview.TextView
	Armor       *tview.TextView
	Trinket     *tview.TextView
	Log         tview.TableContent
}

func NewBase() *Base {
	var b Base
	b.HP = tview.NewTextView()
	b.Mana = tview.NewTextView()
	b.Weapon = tview.NewTextView()
	b.Progression = tview.NewTextView()
	b.Trinket = tview.NewTextView()
	b.HP = tview.NewTextView()
	b.Log = NewActivityLog()

	return &b
}

func BaseFrame() tview.Primitive {
	// build player, monster, and log UI elements
	// p := Player()
	// build a flex here to contain player, monster, and log UI elements
	// frame := tview.NewFrame(p)
	frame := tview.NewFrame(nil)
	border := 1
	text := 2
	frame.SetBorders(border, border, text, text, border, border)

	frame.AddText("Terminal Idler", true, tview.AlignCenter, tcell.ColorRed)
	return frame
}

func Player(p components.Player) tview.Primitive {
	pc := p.PlayerComponent
	hpMeter := tview.NewTextView()
	fmt.Fprintln(hpMeter, TextMeter(pc.Health.Total, pc.Health.Max, 30))

	return hpMeter
}

func TextMeter(a, b, n int) string {
	on := `⊠`
	off := `⊡`
	perc := float64(a) / float64(b)
	bars := int(float64(n) * perc)
	return strings.Repeat(on, bars) + strings.Repeat(off, n-bars)
}
