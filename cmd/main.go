package main

import (
	"os/signal"
	"syscall"
	"time"
	// "fmt"
	"os"

	"github.com/EngoEngine/ecs"

	"github.com/therealfakemoot/ecs-rogue/components"
	"github.com/therealfakemoot/ecs-rogue/systems"
)

func main() {
	w := &ecs.World{}

	// var renderable *systems.Renderable
	// trs := &systems.TerminalRenderSystem{W: os.Stdout, Entities: make(map[uint64]systems.Renderable)}
	// w.AddSystemInterface(trs, renderable, nil)

	var spawnable *systems.Spawnable
	mss := &systems.MobSpawnerSystem{Entities: make(map[uint64]systems.Spawnable)}
	w.AddSystemInterface(mss, spawnable, nil)

	p := components.Player{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Type = components.RenderPlayer
	p.Name = "Jumbo Chungus"
	p.PlayerComponent.Health.Max = 100
	p.PlayerComponent.Health.Total = 13
	p.PlayerComponent.Mana.Max = 100
	p.PlayerComponent.Mana.Total = 69
	// p.Health.Total = 87

	m := components.Mob{BasicEntity: ecs.NewBasic()}
	m.RenderComponent.Type = components.RenderMob

	w.AddEntity(&p)
	// w.AddEntity(&m)
	t := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-t.C:
				w.Update(1)

			}
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	// fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	done <- true
}
