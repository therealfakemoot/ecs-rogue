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
	var spawnable *systems.Spawnable
	trs := &systems.TerminalRenderSystem{W: os.Stdout, Entities: make(map[uint64]systems.Renderable)}
	mss := &systems.MobSpawnerSystem{}
	w.AddSystemInterface(trs, renderable, nil)
	w.AddSystemInterface(mss, spawnable, nil)

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
	logTable := tview.NewTable().SetBorders(true)
	logFrame := tview.NewFrame(logTable)
	NewTableRow(logTable, []string{"level", "You exit the dungeons into frigid mountain winds..."}, 0)
	NewTableRow(logTable, []string{"gold", "+24"}, 1)
	NewTableRow(logTable, []string{"exp", "+5"}, 2)
	NewTableRow(logTable, []string{"item", "morbid finger | sapphire ring"}, 3)
	NewTableRow(logTable, []string{"level", "The mountain foothils turn into rolling plains..."}, 4)
	NewTableRow(logTable, []string{"gold", "+85"}, 5)
	NewTableRow(logTable, []string{"exp", "+24"}, 6)
	logFrame.AddText("Action Log", true, tview.AlignLeft, tcell.ColorGoldenrod)
	logBox := tview.NewFlex().AddItem(logFrame, 0, 1, false)

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
