package main

import (
	// "fmt"
	"os"

	"github.com/EngoEngine/ecs"

	"github.com/therealfakemoot/ecs-rogue/components"
	"github.com/therealfakemoot/ecs-rogue/systems"
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
	w.Update(1)
}
