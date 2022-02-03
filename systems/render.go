package systems

import (
	"io"

	"github.com/EngoEngine/ecs"
)

type TerminalRenderSystem struct {
	W io.Writer
	State string
}

func (trs TerminalRenderSystem) Update(dt float32){
	trs.W.Write([]byte(trs.State))
}

func (trs TerminalRenderSystem) Remove(e ecs.BasicEntity){}
