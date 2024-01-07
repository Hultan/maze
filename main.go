package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

type Rect struct {
	t, l float64
	w, h float64
}

type Dimensions int

const (
	World2D Dimensions = iota
	World3D
)

var maze *mazeGen.Maze
var dim = World3D
var pos = rl.NewVector3(0, 1, -60)
var target = rl.NewVector3(0, 2.5, 0)
var xx, yy = 0, 0
var camera rl.Camera3D

func main() {
	rl.InitWindow(800, 600, "Maze")
	defer rl.CloseWindow()

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	maze = mazeGen.NewMaze()
	rl.DisableCursor() // Limit cursor to relative movement inside the window

	up := rl.NewVector3(0, 1, 0)
	camera = rl.NewCamera3D(pos, target, up, 20, rl.CameraPerspective)

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
		}
	}
}

func draw2D() {
	if xx == 24 && yy == 24 {
		maze = mazeGen.NewMaze()
		xx = 0
		yy = 0
	}
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

	rl.BeginDrawing()

	rl.ClearBackground(rl.DarkGray)

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

	rl.EndDrawing()
}

func draw3D() {
	rl.UpdateCamera(&camera, rl.CameraFree)

	mov := rl.NewVector3(0, 0, 0)
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		mov.X += 0.1
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		mov.X -= 0.1
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		mov.Y += 0.1
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		mov.Y -= 0.1
	}
	rot := rl.NewVector3(rl.GetMouseDelta().X*0.05, rl.GetMouseDelta().Y*0.05, 0)
	rl.UpdateCameraPro(&camera, mov, rot, rl.GetMouseWheelMove()*2)
	//if rl.IsKeyPressed(rl.KeyUp) {
	//	pos.X += 1
	//}
	//if rl.IsKeyPressed(rl.KeyDown) {
	//	pos.X -= 1
	//}
	//if rl.IsKeyPressed(rl.KeyRight) {
	//	pos.Z += 1
	//}
	//if rl.IsKeyPressed(rl.KeyLeft) {
	//	pos.Z -= 1
	//}

	rl.BeginDrawing()

	rl.ClearBackground(rl.White)

	rl.BeginMode3D(camera)

	rl.DrawCube(camera.Target, 0.5, 0.5, 0.5, rl.Purple)
	rl.DrawCubeWires(camera.Target, 0.5, 0.5, 0.5, rl.DarkPurple)

	rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.Beige) // Draw ground
	rl.DrawCube(rl.NewVector3(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)            // Draw a blue wall
	rl.DrawCube(rl.NewVector3(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)             // Draw a green wall
	rl.DrawCube(rl.NewVector3(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)             // Draw a yellow wall

	rl.EndMode3D()
	rl.EndDrawing()
}
