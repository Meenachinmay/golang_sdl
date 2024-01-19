package main

import "github.com/veandco/go-sdl2/sdl"

func NewGame() (*Game, error) {
    // Initialization logic from your original newGame function
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

	game := &Game{
        window: window,
        renderer: renderer,
        running: true,
        players: []*Entity{
            NewEntity(float32(windowWidth)/2, float32(windowHeight)/2, 0, 0, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
        },
		enemies: []*Entity{
            NewEntity(100, 100, 0, 0, sdl.Color{R: 255, G: 0, B: 0, A: 255}), // Red enemy
        },
    }

    return game, nil
}
