package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

var pos = rl.NewVector3(0, 1, -1)
var up = rl.NewVector3(0, 1, 0)
var target = rl.NewVector3(0, 1, 0)
var camera rl.Camera3D

func draw3D() {
	rl.UpdateCamera(&camera, rl.CameraFirstPerson)

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

	rl.BeginDrawing()

	rl.ClearBackground(rl.White)

	rl.BeginMode3D(camera)

	rl.DrawCube(rl.NewVector3(0.0, -0.5, 0.0), 100.0, 1.0, 100.0, rl.Beige) // Draw ground
	rl.DrawCube(rl.NewVector3(0.0, 5.5, 0.0), 100.0, 1.0, 100.0, rl.Brown)  // Draw ground

	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			w := maze[y][x].Walls
			if x == 0 || y == 0 {
				if w&mazeGen.North != 0 {
					rl.DrawCube(rl.NewVector3(coord(x), 2.5, coord(y)), 4.0, 5.0, 1.0, rl.Blue)
					//rl.DrawLine((x+1)*size, (y+1)*size, (x+2)*size, (y+1)*size, rl.RayWhite)
				}
				if w&mazeGen.West != 0 {
					rl.DrawCube(rl.NewVector3(coord(x), 2.5, coord(y+1)), 1.0, 5.0, 4.0, rl.Blue)
					//rl.DrawLine((x+1)*size, (y+1)*size, (x+1)*size, (y+2)*size, rl.RayWhite)
				}
			}
			if w&mazeGen.South != 0 {
				rl.DrawCube(rl.NewVector3(coord(x), 2.5, coord(y+1)), 4.0, 5.0, 1.0, rl.Blue)
				//rl.DrawLine((x+1)*size, (y+2)*size, (x+2)*size, (y+2)*size, rl.RayWhite)
			}
			if w&mazeGen.East != 0 {
				rl.DrawCube(rl.NewVector3(coord(x+1), 2.5, coord(y)), 1.0, 5.0, 4.0, rl.Blue)
				//rl.DrawLine((x+2)*size, (y+1)*size, (x+2)*size, (y+2)*size, rl.RayWhite)
			}
		}
	}

	rl.EndMode3D()
	rl.EndDrawing()
}

func coord(v int) float32 {
	return float32(v)*4 - 50.5
}
