package systems

import (
	"bytes"
	// "fmt"
	"io"
	"log"
	"text/template"

	"github.com/EngoEngine/ecs"
	"github.com/therealfakemoot/ecs-rogue/components"
)

type TerminalRenderSystem struct {
	World    *ecs.World
	W        io.Writer
	Entities map[uint64]Renderable
	T        *template.Template
}

func (trs *TerminalRenderSystem) New(w *ecs.World) {
	trs.World = w
	trs.T = template.Must(template.ParseGlob("templates/*.tpl"))
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
	data := struct {
		Player *components.PlayerComponent
		Mobs   []*components.MobComponent
		MobHP  int
	}{}
	for _, e := range trs.Entities {
		v := e.(components.RenderComponentInterface)
		rc := v.GetRenderComponent()
		switch rc.Type {
		case components.RenderPlayer:
			p := e.(components.PlayerComponentInterface)
			pc := p.GetPlayerComponent()
			data.Player = pc
			// fmt.Fprintf(&b, "%s(hp/hpMax)(mana/manaMax)\n", pc.Name)
		case components.RenderMob:
			m := e.(components.MobComponentInterface)
			mob := m.GetMobComponent()
			data.MobHP += mob.Health.Max
			data.Mobs = append(data.Mobs, mob)
		}
	}
	err := trs.T.ExecuteTemplate(&b, "player_tick.tpl", data)
	if err != nil {
		log.Fatalf("error rendering player template: %s", err)
	}
	err = trs.T.ExecuteTemplate(&b, "mob_tick.tpl", data)
	if err != nil {
		log.Fatalf("error rendering mob template: %s", err)
	}
	io.Copy(trs.W, &b)
}

func (trs *TerminalRenderSystem) Remove(e ecs.BasicEntity) {
	delete(trs.Entities, e.ID())
}
