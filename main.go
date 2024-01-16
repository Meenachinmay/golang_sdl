package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
	boxSize      = 50
	targetFPS    = 60
	boxSpeed     = 1000 // pixels per second
)

type game struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	running  bool
	boxX     float32
	boxY     float32
	velX     float32
	velY     float32
}

func newGame() (*game, error) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		window.Destroy()
		return nil, err
	}

	return &game{window: window, renderer: renderer, running: true, boxX: float32(windowWidth) / 2, boxY: float32(windowHeight) / 2}, nil

}

func (g *game) run() {
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

func (g *game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			g.running = false
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				switch e.Keysym.Sym {
				case sdl.K_w, sdl.K_UP:
					g.velY = -1
				case sdl.K_s, sdl.K_DOWN:
					g.velY = 1
				case sdl.K_a, sdl.K_LEFT:
					g.velX = -1
				case sdl.K_d, sdl.K_RIGHT:
					g.velX = 1
				case sdl.K_ESCAPE:
					g.running = false
					sdl.Quit()
				}
			} else if e.Type == sdl.KEYUP {
				switch e.Keysym.Sym {
                case sdl.K_w, sdl.K_s, sdl.K_UP, sdl.K_DOWN:
                    g.velY = 0
                case sdl.K_a, sdl.K_d, sdl.K_LEFT, sdl.K_RIGHT:
                    g.velX = 0
                }
			}
		}
	}
}

func (g *game) update(deltaTime float32) {
    g.boxX += g.velX * boxSpeed * deltaTime
    g.boxY += g.velY * boxSpeed * deltaTime
}


func (g *game) render() {
	g.renderer.SetDrawColor(0, 0, 0, 255) // Black background
	g.renderer.Clear()

	g.renderer.SetDrawColor(255, 255, 255, 255) // White box
	g.renderer.FillRect(&sdl.Rect{X: int32(g.boxX), Y: int32(g.boxY), W: boxSize, H: boxSize})

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
