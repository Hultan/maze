package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

const compassSize, compassX, compassY = 70, 700, 100

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

	rl.ClearBackground(rl.White)

	rl.BeginMode3D(camera)

	const size = 25
	const width, thick, height = 5, 1, 3
	var wallColor = rl.DarkBrown

	rl.DrawCube(rl.NewVector3(0.0, -thick, 0.0), 4*size+1, thick, 4*size+1, rl.Beige) // Draw ground
	rl.DrawCube(rl.NewVector3(0.0, height, 0.0), 4*size+1, thick, 4*size+1, rl.Brown) // Draw roof

	start := rl.NewVector3(-size*2+2, height/2, -size*2+2)
	end := rl.NewVector3(size*2-2, height/2, size*2-2)
	rl.DrawCube(start, thick, thick, thick, rl.Green) // Draw start cube
	rl.DrawCube(end, thick, thick, thick, rl.Blue)    // Draw end cube

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
	rl.DrawCircleLines(compassX, compassY, compassSize, rl.Red)
	DrawNeedle(camera.Target, camera.Position, rl.Black)
	DrawNeedle(end, camera.Position, rl.Blue)

	// Map
	xx = int((camera.Position.X + size*2) / 4)
	yy = int((camera.Position.Z + size*2) / 4)

	if rl.IsKeyDown(rl.KeyM) {
		draw2D(false)
	}
}

func DrawNeedle(start, end rl.Vector3, col color.RGBA) {
	cx := start.X - end.X
	cy := start.Z - end.Z
	cx, cy = normalize(cx, cy, compassSize)
	rl.DrawLine(compassX, compassY, int32(compassX+cx), int32(compassY+cy), col)
}

func normalize(cx, cy, size float32) (float32, float32) {
	l := float32(math.Sqrt(float64(cx*cx + cy*cy)))
	return cx / l * size, cy / l * size
}

func coord(v int) float32 {
	return float32(v)*4 - 50
}
