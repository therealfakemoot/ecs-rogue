package systems

import (
	"github.com/EngoEngine/ecs"
	// "github.com/therealfakemoot/ecs-rogue/components"
)

type MobSpawnerSystem struct {
	World *ecs.World
}

func (mss *MobSpawnerSystem) New(w *ecs.World) {
	mss.World = w
}

func (mss *MobSpawnerSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Renderable)
	mss.Add(obj)
}

func (mss *MobSpawnerSystem) Add(r Renderable) {
	// mss.Entities[r.GetBasicEntity().ID()] = r
	// mss.Entities = append(mss.Entities, r)
}

func (mss MobSpawnerSystem) Update(dt float32)        {}
func (mss MobSpawnerSystem) Remove(e ecs.BasicEntity) {}
