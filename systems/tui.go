package systems

import (
	// "bytes"
	// "fmt"
	// "io"
	// "log"
	"text/template"

	"github.com/EngoEngine/ecs"
	"github.com/rivo/tview"
	"github.com/therealfakemoot/ecs-rogue/components"
)

type TUIRenderSystem struct {
	App      *tview.Application
	World    *ecs.World
	Entities map[uint64]Renderable
	T        *template.Template
}

func (trs *TUIRenderSystem) New(w *ecs.World) {
	trs.World = w
	trs.T = template.Must(template.ParseGlob("templates/*.tpl"))
}

func (trs *TUIRenderSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Renderable)
	trs.Add(obj)
}

func (trs *TUIRenderSystem) Add(r Renderable) {
	trs.Entities[r.GetBasicEntity().ID()] = r
	// trs.Entities = append(trs.Entities, r)
}

func (trs TUIRenderSystem) Update(dt float32) {
	data := struct {
		Player *components.PlayerComponent
		Mobs   []*components.MobComponent
		MobHP  int
	}{}
	// first step is to create a statemap
	// players and mobs are the main thing being rendered, but i'll need to make sure
	// to cover things like screen effects:
	//   writing garbage bytes to a random subset of on-screen locations to distort the UI
	//   ??????

	// HP Bar TextView
	// Mana Bar TextView
	// Weapon View
	// Armor View
	// Trinket View
	go func() {
		trs.App.QueueUpdateDraw(func() {
		})
	}()
	for _, e := range trs.Entities {
		v := e.(components.RenderComponentInterface)
		rc := v.GetRenderComponent()
		switch rc.Type {
		case components.RenderPlayer:
			p := e.(components.PlayerComponentInterface)
			pc := p.GetPlayerComponent()
			data.Player = pc
		case components.RenderMob:
			m := e.(components.MobComponentInterface)
			mob := m.GetMobComponent()
			data.MobHP += mob.Health.Max
			data.Mobs = append(data.Mobs, mob)
		}
	}
}

func (trs *TUIRenderSystem) Remove(e ecs.BasicEntity) {
	delete(trs.Entities, e.ID())
}
