package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

type Dimensions int

const (
	World2D Dimensions = iota
	World3D
)

var maze *mazeGen.Maze
var dim = World2D
var startingPos = rl.NewVector3(2, 1, 2)
var exitWindow = false

func main() {
	rl.InitWindow(800, 600, "Maze")
	defer rl.CloseWindow()
	rl.SetExitKey(rl.KeyNull)

	rl.MaximizeWindow()
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)
	rl.DisableCursor() // Limit cursor to relative movement inside the window
	camera = rl.NewCamera3D(startingPos, target, up, 60, rl.CameraPerspective)

	for !exitWindow {
		if rl.IsKeyDown(rl.KeyTwo) {
			dim = World2D
		}
		if rl.IsKeyDown(rl.KeyThree) {
			dim = World3D
		}
		if rl.IsKeyDown(rl.KeyEscape) || rl.IsKeyDown(rl.KeyQ) {
			exitWindow = true
			return
		}

		if maze == nil || xx == maze.Width()-1 && yy == maze.Height()-1 {
			maze = mazeGen.NewMaze(30, 20)
			xx = 0
			yy = 0
			camera.Position = startingPos
		}

		if !maze.IsDone() {
			maze.Generate()
		} else {
			if maze.SolutionStarted == false {
				maze.SolutionStarted = true
				maze.StartSolving()
			}
			time.Sleep(time.Millisecond * 100)
			maze.SolveStep()
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
