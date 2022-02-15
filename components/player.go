package components

import (
	"github.com/EngoEngine/ecs"
)

type Player struct {
	ecs.BasicEntity
	PlayerComponent
	RenderComponent
	MobComponent
}
type PlayerComponentInterface interface {
	GetPlayerComponent() *PlayerComponent
}

type PlayerComponent struct {
	Name string
}

func (pc *PlayerComponent) GetPlayerComponent() *PlayerComponent {
	return pc
}
