package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

var xx, yy = 0, 0

func draw2D(clear bool) {
	wall := maze[yy][xx].Walls
	if rl.IsKeyPressed(rl.KeyUp) {
		if wall&mazeGen.North == 0 {
			yy -= 1
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		if wall&mazeGen.South == 0 {
			yy += 1
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		if wall&mazeGen.West == 0 {
			xx -= 1
		}
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		if wall&mazeGen.East == 0 {
			xx += 1
		}
	}

	if clear {
		rl.ClearBackground(rl.DarkGray)
	}

	size := int32(600 / 30)

	for x := int32(0); x < 25; x++ {
		for y := int32(0); y < 25; y++ {
			w := maze[y][x].Walls
			if w&mazeGen.North != 0 {
				rl.DrawLine((x+1)*size, (y+1)*size, (x+2)*size, (y+1)*size, rl.RayWhite)
			}
			if w&mazeGen.East != 0 {
				rl.DrawLine((x+2)*size, (y+1)*size, (x+2)*size, (y+2)*size, rl.RayWhite)
			}
			if w&mazeGen.South != 0 {
				rl.DrawLine((x+1)*size, (y+2)*size, (x+2)*size, (y+2)*size, rl.RayWhite)
			}
			if w&mazeGen.West != 0 {
				rl.DrawLine((x+1)*size, (y+1)*size, (x+1)*size, (y+2)*size, rl.RayWhite)
			}
		}
	}

	// Draw green starting position
	rl.DrawRectangle(size+3, size+3, size-6, size-6, rl.Green)

	// Draw blue starting position
	rl.DrawRectangle(size*25+3, size*25+3, size-6, size-6, rl.Blue)

	// TODO : Draw yellow player position
	rl.DrawRectangle(size*(int32(xx)+1)+3, size*(int32(yy)+1)+3, size-6, size-6, rl.Yellow)
}
