package components

import (
	"github.com/EngoEngine/ecs"
)

type DamageType int

const (
	PhysicalDamage DamageType = iota
	FireDamage
	WaterDamage
	EarthDamage
	AirDamage
	ElectricDamage
	PsychicDamage
	NecroticDamage
	RadiantDamage
	EntropicDamage
)

type PackModifier struct {
	HPFlat      int     // directly added to HP
	HPMult      float64 // multiplied
	ManaFlat    int
	ManaMult    float64
	Resistances map[DamageType]float64
	Buffs       map[string]func() // the signature of this func is really critical but i have to skeleton out the rest of this dinkle doink to find out
}

type Pack struct {
	ecs.BasicEntity
	RenderComponent
}

type PackComponent struct {
	Modifiers []PackModifier
	Members   []ecs.IdentifierSlice
}
