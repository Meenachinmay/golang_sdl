// package main

// import "github.com/veandco/go-sdl2/sdl"

// func main() {
// 	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
// 		panic(err)
// 	}
// 	defer sdl.Quit()

// 	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
// 		800, 600, sdl.WINDOW_SHOWN)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer window.Destroy()

// 	surface, err := window.GetSurface()
// 	if err != nil {
// 		panic(err)
// 	}
// 	surface.FillRect(nil, 0)

// 	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
// 	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
// 	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
// 	surface.FillRect(&rect, pixel)
// 	window.UpdateSurface()

// 	running := true
// 	for running {
// 		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
// 			switch event.(type) {
// 			case *sdl.QuitEvent:
// 				println("Quit")
// 				running = false
// 				break
// 			}
// 		}
// 	}
// }

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
	boxSize      = 50
)

type game struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	running  bool
	boxX     int32
	boxY     int32
}

func newGame() (*game, error) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	return &game{window: window, renderer: renderer, running: true, boxX: windowWidth / 2, boxY: windowHeight / 2}, nil
}

func (g *game) run() {
	for g.running {
		g.handleEvents()
		g.update()
		g.render()
		sdl.Delay(16) // Roughly 60 frames per second
	}

	g.cleanup()
}

func (g *game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			g.running = false
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				switch e.Keysym.Sym {
				case sdl.K_w:
					g.boxY -= 10
				case sdl.K_s:
					g.boxY += 10
				case sdl.K_a:
					g.boxX -= 10
				case sdl.K_d:
					g.boxX += 10
				}
			}
		}
	}
}

func (g *game) update() {
	// Any game logic updates would go here
}

func (g *game) render() {
	g.renderer.SetDrawColor(0, 0, 0, 255) // Black background
	g.renderer.Clear()

	g.renderer.SetDrawColor(255, 255, 255, 255) // White box
	g.renderer.FillRect(&sdl.Rect{g.boxX, g.boxY, boxSize, boxSize})

	g.renderer.Present()
}

func (g *game) cleanup() {
	g.window.Destroy()
	g.renderer.Destroy()
	sdl.Quit()
}

func main() {
	g, err := newGame()
	if err != nil {
		panic(err)
	}
	g.run()
}
