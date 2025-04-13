package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

var xx, yy int32 = 0, 0

func draw2D(clear bool) {
	wall := maze.Cells[yy][xx].Walls
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

	// TODO : Fix size calculation
	size := int32(600 / 30)

	for x := int32(0); x < maze.Width(); x++ {
		for y := int32(0); y < maze.Height(); y++ {
			w := maze.Cells[y][x].Walls
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

	// Draw blue ending position
	rl.DrawRectangle(size*maze.Width()+3, size*maze.Height()+3, size-6, size-6, rl.Blue)

	// Draw yellow current position
	rl.DrawRectangle(size*(int32(xx)+1)+3, size*(int32(yy)+1)+3, size-6, size-6, rl.Yellow)

	for i, step := range maze.SolutionStack {
		var col = rl.Red
		if i == len(maze.SolutionStack)-1 {
			col = rl.Violet
		}
		rl.DrawRectangle(size*(step.From.X+1)+3, size*(step.From.Y+1)+3, size-6, size-6, col)
	}

	if maze.SolutionDone {
		for _, step := range maze.SolutionPath {
			rl.DrawRectangle(size*(step.X+1)+3, size*(step.Y+1)+3, size-6, size-6, rl.Violet)
		}
	}
}
