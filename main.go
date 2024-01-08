package main

import (
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

func main() {
	rl.InitWindow(800, 600, "Maze")
	defer rl.CloseWindow()

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	maze = mazeGen.NewMaze()
	rl.DisableCursor() // Limit cursor to relative movement inside the window

	camera = rl.NewCamera3D(pos, target, up, 60, rl.CameraPerspective)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyTwo) {
			dim = World2D
		}
		if rl.IsKeyDown(rl.KeyThree) {
			dim = World3D
		}
		switch dim {
		case World2D:
			draw2D()
		case World3D:
			draw3D()
			//draw2D()
		}
	}
}
