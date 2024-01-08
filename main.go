package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

type Dimensions int

const (
	World2D Dimensions = iota
	World3D
)

var maze *mazeGen.Maze
var dim = World3D
var startingPos = rl.NewVector3(-48, 1, -48)

func main() {
	rl.InitWindow(800, 600, "Maze")
	defer rl.CloseWindow()

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	maze = mazeGen.NewMaze()
	rl.DisableCursor() // Limit cursor to relative movement inside the window

	camera = rl.NewCamera3D(startingPos, target, up, 60, rl.CameraPerspective)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyTwo) {
			dim = World2D
		}
		if rl.IsKeyDown(rl.KeyThree) {
			dim = World3D
		}
		fmt.Println(xx, yy)
		if xx == 24 && yy == 24 {
			maze = mazeGen.NewMaze()
			xx = 0
			yy = 0
			camera.Position = startingPos
		}

		rl.BeginDrawing()

		switch dim {
		case World2D:
			draw2D(true)
		case World3D:
			draw3D()
		}

		rl.EndDrawing()
	}
}
