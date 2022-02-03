package main

import (
	"github.com/EngoEngine/ecs"

	"github.com/therealfakemoot/ecs-rogue/systems"
)

func main() {
	w := ecs.World{}

	w.AddSystem(systems.TerminalRenderSystem{})
}
