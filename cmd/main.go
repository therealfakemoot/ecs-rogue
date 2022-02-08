package main

import (
	"os"

	"github.com/EngoEngine/ecs"

	"github.com/therealfakemoot/ecs-rogue/components"
	"github.com/therealfakemoot/ecs-rogue/systems"
)

func main() {
	w := &ecs.World{}

	var renderable *systems.Renderable
	trs := &systems.TerminalRenderSystem{W: os.Stdout}
	w.AddSystemInterface(trs, renderable, nil)

	p := components.Player{}
	p.RenderComponent.Type = components.RenderPlayer
	p.Name = "Jumbo Chungus"

	m := components.Mob{}

	w.AddEntity(&p)
	w.AddEntity(&m)
	w.Update(1)
}
