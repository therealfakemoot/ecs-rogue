package components

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/fatih/color"
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
	Name    string
	MinHP   int
	MaxHP   int
	MinMana int
	MaxMana int
}

func (pc *PlayerComponent) GetPlayerComponent() *PlayerComponent {
	return pc
}

func (pc PlayerComponent) Health() string {
	d := float64(pc.MinHP) / float64(pc.MaxHP)
	switch {
	case d < .25:
		return color.RedString(fmt.Sprintf("%d/%d", pc.MinHP, pc.MaxHP))
	case d < .50:
		return color.YellowString(fmt.Sprintf("%d/%d", pc.MinHP, pc.MaxHP))
	default:
		return color.GreenString(fmt.Sprintf("%d/%d", pc.MinHP, pc.MaxHP))
	}
}

func (pc PlayerComponent) Mana() string {
	d := float64(pc.MinMana) / float64(pc.MaxMana)
	switch {
	case d < .25:
		return color.RedString(fmt.Sprintf("%d/%d", pc.MinMana, pc.MaxMana))
	case d < .50:
		return color.YellowString(fmt.Sprintf("%d/%d", pc.MinMana, pc.MaxMana))
	default:
		return color.BlueString(fmt.Sprintf("%d/%d", pc.MinMana, pc.MaxMana))
	}
}
