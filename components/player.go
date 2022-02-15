package components

import (
	// "fmt"

	"github.com/EngoEngine/ecs"
	// "github.com/fatih/color"
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
	Name   string
	Health ResourceBarComponent
	Mana   ResourceBarComponent
}

func (pc *PlayerComponent) GetPlayerComponent() *PlayerComponent {
	return pc
}
