package systems

import (
	"bytes"
	"fmt"
	"io"
	// "log"

	"github.com/EngoEngine/ecs"
	"github.com/therealfakemoot/ecs-rogue/components"
)

type TerminalRenderSystem struct {
	World    *ecs.World
	W        io.Writer
	Entities map[uint64]Renderable
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
	trs.Add(obj)
}

func (trs *TerminalRenderSystem) Add(r Renderable) {
	trs.Entities[r.GetBasicEntity().ID()] = r
	// trs.Entities = append(trs.Entities, r)
}

func (trs TerminalRenderSystem) Update(dt float32) {
	var b bytes.Buffer
	for _, e := range trs.Entities {
		// log.Printf("rendering entity %d: %#+v\n", id, e)
		v := e.(components.RenderComponentInterface)
		rc := v.GetRenderComponent()
		switch rc.Type {
		case components.RenderPlayer:
			p := e.(components.PlayerComponentInterface)
			pc := p.GetPlayerComponent()
			fmt.Fprintf(&b, "%s(hp/hpMax)(mana/manaMax)\n", pc.Name)

		}
	}
	io.Copy(trs.W, &b)
}

func (trs *TerminalRenderSystem) Remove(e ecs.BasicEntity) {
	delete(trs.Entities, e.ID())
}
