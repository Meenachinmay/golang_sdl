// update.go
package main

func (g *Game) update(deltaTime float32) {
    g.boxX += g.velX * boxSpeed * deltaTime
    g.boxY += g.velY * boxSpeed * deltaTime
}