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
    players []*Entity
    enemies []*Entity
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

func (g *Game) update(deltaTime float32) {
    for _, player := range g.players {
        player.Update(deltaTime)
    }
    for _, enemy := range g.enemies {
        enemy.Update(deltaTime)
    }
    // Additional game update logic
}

func (g *Game) render() {
    g.renderer.SetDrawColor(0, 0, 0, 255) // Black background
    g.renderer.Clear()

    for _, player := range g.players {
        player.Render(g.renderer)
    }
    for _, enemy := range g.enemies {
        enemy.Render(g.renderer)
    }

    g.renderer.Present()
}

func (g *Game) cleanup() {
	g.window.Destroy()
	g.renderer.Destroy()
	sdl.Quit()
}