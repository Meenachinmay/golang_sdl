// game.go
package main

import (
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
    window   *sdl.Window
    renderer *sdl.Renderer
    running  bool
    boxX     float32
    boxY     float32
    velX     float32
    velY     float32
}

// Method definitions for Game (e.g., run, cleanup) go here
func (g *Game) run() {
    var lastUpdate time.Time = time.Now()

    for g.running {
        currentTime := time.Now()
        deltaTime := currentTime.Sub(lastUpdate).Seconds() // Delta time in seconds
        lastUpdate = currentTime

        g.handleEvents()
        g.update(float32(deltaTime))
        g.render()
    }

    g.cleanup()
}

func (g *Game) cleanup() {
	g.window.Destroy()
	g.renderer.Destroy()
	sdl.Quit()
}