// events.go
package main

import "github.com/veandco/go-sdl2/sdl"

func (g *Game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			g.running = false
		case *sdl.KeyboardEvent:
			player := g.players[0]
			if e.Type == sdl.KEYDOWN {
				switch e.Keysym.Sym {
				case sdl.K_w, sdl.K_UP:
					player.VelY = -1
				case sdl.K_s, sdl.K_DOWN:
					player.VelY = 1
				case sdl.K_a, sdl.K_LEFT:
					player.VelX = -1
				case sdl.K_d, sdl.K_RIGHT:
					player.VelX = 1
				case sdl.K_ESCAPE:
					g.running = false
					sdl.Quit()
				}
			} else if e.Type == sdl.KEYUP {
				switch e.Keysym.Sym {
                case sdl.K_w, sdl.K_s, sdl.K_UP, sdl.K_DOWN:
                    player.VelY = 0
                case sdl.K_a, sdl.K_d, sdl.K_LEFT, sdl.K_RIGHT:
                    player.VelX = 0
                }
			}
		}
	}
}