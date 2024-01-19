// entity.go
package main

import "github.com/veandco/go-sdl2/sdl"

type Entity struct {
	X, Y       float32
	VelX, VelY float32
	Color      sdl.Color
	Texture    *sdl.Texture
	Angle      float64
}

func NewEntity(x, y, velX, velY float32, color sdl.Color) *Entity {
	return &Entity{
		X:     x,
		Y:     y,
		VelX:  velX,
		VelY:  velY,
		Color: color,
	}
}

func (e *Entity) Update(deltaTime float32) {
	e.X += e.VelX * boxSpeed * deltaTime
	e.Y += e.VelY * boxSpeed * deltaTime
}

func (e *Entity) Render(renderer *sdl.Renderer) {
	rect := &sdl.Rect{X: int32(e.X), Y: int32(e.Y), W: boxSize, H: boxSize}
	renderer.CopyEx(e.Texture, nil, rect, e.Angle, nil, sdl.FLIP_NONE)
}
