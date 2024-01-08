package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

var pos = rl.NewVector3(-25*2+2, 1, -25*2+2)
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

	const size = 25
	const width, thick, height = 5, 1, 3
	var wallColor = rl.DarkBrown

	rl.DrawCube(rl.NewVector3(0.0, -thick, 0.0), 4*size+1, thick, 4*size+1, rl.Beige) // Draw ground
	rl.DrawCube(rl.NewVector3(0.0, height, 0.0), 4*size+1, thick, 4*size+1, rl.Brown) // Draw roof

	rl.DrawCube(rl.NewVector3(-size*2+2, height/2, -size*2+2), thick, thick, thick, rl.Green) // Draw start cube
	rl.DrawCube(rl.NewVector3(size*2-2, height/2, size*2-2), thick, thick, thick, rl.Blue)    // Draw end cube

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			w := maze[y][x].Walls
			if x == 0 || y == 0 {
				if w&mazeGen.North != 0 {
					rl.DrawCube(rl.NewVector3(coord(x)+width/2, height/2, coord(y)), width, height, thick, wallColor)
				}
				if w&mazeGen.West != 0 {
					rl.DrawCube(rl.NewVector3(coord(x), height/2, coord(y)+width/2), thick, height, width, wallColor)
				}
			}
			if w&mazeGen.South != 0 {
				rl.DrawCube(rl.NewVector3(coord(x)+width/2, height/2, coord(y+1)), width, height, thick, wallColor)
			}
			if w&mazeGen.East != 0 {
				rl.DrawCube(rl.NewVector3(coord(x+1), height/2, coord(y)+width/2), thick, height, width, wallColor)
			}
		}
	}

	rl.EndMode3D()

	// Compass

	rl.EndDrawing()
}

func coord(v int) float32 {
	return float32(v)*4 - 50
}
