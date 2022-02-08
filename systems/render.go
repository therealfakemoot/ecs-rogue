package systems

import (
	"fmt"
	"io"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/therealfakemoot/ecs-rogue/components"
)

type TerminalRenderSystem struct {
	World    *ecs.World
	W        io.Writer
	Entities ecs.IdentifierSlice
}

func (trs *TerminalRenderSystem) New(w *ecs.World) {
	trs.World = w
}

type Renderable interface {
	ecs.BasicFace
	components.RenderComponentInterface
}

func (trs *TerminalRenderSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Renderable)
	trs.Add(obj.GetBasicEntity())
}

func (trs *TerminalRenderSystem) Add(b *ecs.BasicEntity) {
	trs.Entities = append(trs.Entities, b)
}

func (trs TerminalRenderSystem) Update(dt float32) {
	for _, e := range trs.Entities {
		v, ok := e.(components.RenderComponentInterface)
		if !ok {
			log.Fatalf("shit 1")
		}
		rc := v.GetRenderComponent()
		switch rc.Type {
		case components.RenderPlayer:
			p, ok := e.(components.PlayerComponentInterface)
			if !ok {
				log.Fatalf("shit 1")
			}
			pc := p.GetPlayerComponent()
			fmt.Fprintf(trs.W, "%s(100/100)(45/75)", pc.Name)

		}
		log.Printf("%#+v\n", e)
	}
}

func (trs *TerminalRenderSystem) Remove(e ecs.BasicEntity) {}
