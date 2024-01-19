package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func NewGame() (*Game, error) {
	// Initialization logic from your original newGame function
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		window.Destroy()
		return nil, err
	}

	// After creating the renderer
	playerTexture, err := createPlayerTexture(renderer)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// Enemy texture
	enemyTexture, err := createEntityTexture(renderer, sdl.Color{R: 255, G: 0, B: 0, A: 255}) // Red color for enemy
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	game := &Game{
		window:   window,
		renderer: renderer,
		running:  true,
		players: []*Entity{
			{X: float32(windowWidth) / 2, Y: float32(windowHeight) / 2, Texture: playerTexture},
		},
		enemies: []*Entity{
			{X: 100, Y: 100, Texture: enemyTexture},
		},
	}

	return game, nil
}

// create player texture
func createPlayerTexture(renderer *sdl.Renderer) (*sdl.Texture, error) {
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_TARGET, boxSize, boxSize)
	if err != nil {
		return nil, err
	}

	renderer.SetRenderTarget(texture)

	// Draw white rectangle
	renderer.SetDrawColor(255, 255, 255, 255) // White
	renderer.Clear()

	// Draw black arrow
	renderer.SetDrawColor(0, 0, 0, 255) // Black
	// Define the points for the arrow triangle
	arrow := []sdl.Point{
		{X: boxSize / 2, Y: 5},             // Top point
		{X: boxSize/2 - 5, Y: boxSize - 5}, // Bottom left
		{X: boxSize/2 + 5, Y: boxSize - 5}, // Bottom right
	}
	renderer.DrawLines(arrow)

	renderer.SetRenderTarget(nil)

	return texture, nil
}

// create general entity
func createEntityTexture(renderer *sdl.Renderer, color sdl.Color) (*sdl.Texture, error) {
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_TARGET, boxSize, boxSize)
	if err != nil {
		return nil, err
	}

	renderer.SetRenderTarget(texture)
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.Clear()
	renderer.SetRenderTarget(nil)

	return texture, nil
}
