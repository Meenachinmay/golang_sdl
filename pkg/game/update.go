// update.go
package main

import (
    "github.com/meenachinmay/golang_sdl/pkg/utils"
)

func (g *Game) update(deltaTime float32) {
    g.boxX += g.velX * boxSpeed * deltaTime
    g.boxY += g.velY * boxSpeed * deltaTime
}