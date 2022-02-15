package components

import (
// "github.com/EngoEngine/ecs"
)

type RenderComponentInterface interface {
	GetRenderComponent() *RenderComponent
}

type RenderComponent struct {
	Vis     bool
	Type    RenderType
	Payload map[string]string
}

func (rc *RenderComponent) GetRenderComponent() *RenderComponent {
	return rc
}
