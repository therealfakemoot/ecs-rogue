package components

import (
	"github.com/EngoEngine/ecs"
)

type RenderType int

const (
	RenderPlayer RenderType = iota
	RenderMob
	RenderItem
	RenderScene
	RenderTransition
)

type Mob struct {
	ecs.BasicEntity
	RenderComponent
	MobComponent
}

type MobComponentInterface interface {
	GetMobComponent() *MobComponent
}

type MobComponent struct {
	Health ResourceBarComponent
	SpatialComponent
}

func (mc *MobComponent) GetMobComponent() *MobComponent {
	return mc
}

type SpatialComponentInterface interface {
	GetSpatialComponent() *SpatialComponent
}

type SpatialComponent struct {
	X int
	Y int
}

func (sc *SpatialComponent) SpatialComponent() *SpatialComponent {
	return sc
}
