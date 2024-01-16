package main

import "github.com/veandco/go-sdl2/sdl"

func (g *Game) render() {
	g.renderer.SetDrawColor(0, 0, 0, 255) // Black background
	g.renderer.Clear()

	g.renderer.SetDrawColor(255, 255, 255, 255) // White box
	g.renderer.FillRect(&sdl.Rect{X: int32(g.boxX), Y: int32(g.boxY), W: boxSize, H: boxSize})

	g.renderer.Present()
}