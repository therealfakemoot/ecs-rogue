package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/therealfakemoot/ecs-rogue/components"
)

type Spawnable interface {
	ecs.BasicFace
	components.MobComponentInterface
}

type MobSpawnerSystem struct {
	World    *ecs.World
	Entities map[uint64]Spawnable
	Clock    float64
}

func (mss *MobSpawnerSystem) New(w *ecs.World) {
	mss.World = w
}

func (mss *MobSpawnerSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Spawnable)
	mss.Add(obj)
}

func (mss *MobSpawnerSystem) Add(s Spawnable) {
	mss.Entities[s.GetBasicEntity().ID()] = s
	// mss.Entities = append(mss.Entities, r)
}

func (mss MobSpawnerSystem) Update(dt float32) {
	// here's where i check how many enemies there are, how long we've been playing, the difficult level, etc and then spawn more
	// i want to use the update interval as part of determining spawn rate but i'm not sure how to build a formula for it
	m := NewMob(1, 0)
	mss.World.AddEntity(&m)
	mss.Clock += float64(dt)

}

func (mss MobSpawnerSystem) Remove(e ecs.BasicEntity) {}

// NewMob generates a single new Mob unit.
//   level: the starting level of the generated mob. this determines starting HP, Mana, how many and which modifiers could be applied, etc
//   mod: 0 disables modifiers; higher values make modifiers more likely, more plentiful, and stronger
func NewMob(level int, mod int) components.Mob {
	m := components.Mob{BasicEntity: ecs.NewBasic()}
	m.RenderComponent.Type = components.RenderMob
	m.Health.Max = level * 30
	m.Health.Total = m.Health.Max

	return m
}
