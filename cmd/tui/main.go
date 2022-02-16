package main

import (
	// "fmt"
	"os"

	"github.com/EngoEngine/ecs"

	"github.com/therealfakemoot/ecs-rogue/components"
	"github.com/therealfakemoot/ecs-rogue/systems"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	w := &ecs.World{}

	var renderable *systems.Renderable
	trs := &systems.TerminalRenderSystem{W: os.Stdout, Entities: make(map[uint64]systems.Renderable)}
	w.AddSystemInterface(trs, renderable, nil)

	p := components.Player{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Type = components.RenderPlayer
	p.Name = "Jumbo Chungus"
	p.PlayerComponent.Health.Max = 100
	p.PlayerComponent.Health.Total = 13
	p.PlayerComponent.Mana.Max = 100
	p.PlayerComponent.Mana.Total = 69
	// p.Health.Total = 87

	m := components.Mob{BasicEntity: ecs.NewBasic()}
	m.RenderComponent.Type = components.RenderMob

	w.AddEntity(&p)
	w.AddEntity(&m)
	// w.Update(1)

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)
	NewTableRow(table, []string{"level", "You exit the dungeons into frigid mountain winds..."}, 0)
	NewTableRow(table, []string{"gold", "+24"}, 1)
	NewTableRow(table, []string{"exp", "+5"}, 2)
	NewTableRow(table, []string{"item", "morbid finger | sapphire ring"}, 3)
	NewTableRow(table, []string{"level", "The mountain foothils turn into rolling plains..."}, 4)
	// logBox := tview.NewBox().SetBorder(true).SetTitle("Log")
	logBox := tview.NewFlex().AddItem(table, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(logBox, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Player"), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Enemies"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Progression"), 5, 1, false), 0, 2, false)
	frame := tview.NewFrame(flex)
	border := 1
	text := 2
	frame.SetBorders(border, border, text, text, border, border)

	frame.AddText("Terminal Idler", true, tview.AlignCenter, tcell.ColorRed)

	if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
		panic(err)
	}
	return
}
