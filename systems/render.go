package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/therealfakemoot/ecs-rogue/components"
)

type Renderable interface {
	ecs.BasicFace
	components.RenderComponentInterface
}
