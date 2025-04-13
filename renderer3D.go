package main

import (
	"fmt"
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hultan/maze/mazeGen"
)

const compassSize, compassX, compassY = 70, 700, 100

var up = rl.NewVector3(0, 1, 0)
var target = rl.NewVector3(0, 1, 0)
var camera rl.Camera3D
var walls []rl.Rectangle
var oldCameraPosition rl.Vector3

func draw3D() {
	walls = []rl.Rectangle{}
	oldCameraPosition = camera.Position
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

	const width, thick, height = 5, 1, 3
	var wallColor = rl.DarkBrown

	rl.DrawCube(rl.NewVector3(50.0, -thick, 50.0), 101, thick, 101, rl.Beige)   // Draw ground
	rl.DrawCube(rl.NewVector3(50.0, height+1, 50.0), 101, thick, 101, rl.Brown) // Draw roof

	start := rl.NewVector3(2, height/2, 2)
	end := rl.NewVector3(98, height/2, 98)
	rl.DrawCube(start, thick, thick, thick, rl.Green) // Draw start cube
	rl.DrawCube(end, thick, thick, thick, rl.Blue)    // Draw end cube

	for x := int32(0); x < maze.Width(); x++ {
		for y := int32(0); y < maze.Height(); y++ {
			w := maze.Cells[y][x].Walls
			if x == 0 || y == 0 {
				if w&mazeGen.North != 0 {
					drawWall(coord(x)+width/2, height/2, coord(y), width, height, thick, wallColor)
				}
				if w&mazeGen.West != 0 {
					drawWall(coord(x), height/2, coord(y)+width/2, thick, height, width, wallColor)
				}
			}
			if w&mazeGen.South != 0 {
				drawWall(coord(x)+width/2, height/2, coord(y+1), width, height, thick, wallColor)
			}
			if w&mazeGen.East != 0 {
				drawWall(coord(x+1), height/2, coord(y)+width/2, thick, height, width, wallColor)
			}
		}
	}
	rl.EndMode3D()

	// Compass
	rl.DrawCircleLines(compassX, compassY, compassSize, rl.Red)
	drawNeedle(camera.Target, camera.Position, rl.Black)
	drawNeedle(end, camera.Position, rl.Blue)

	// Map
	xx = int32((camera.Position.X + float32(maze.Width())*2) / 4)
	yy = int32((camera.Position.Z + float32(maze.Height())*2) / 4)
	if rl.IsKeyDown(rl.KeyM) {
		draw2D(false)
	}

	// Draw position
	s := fmt.Sprintf("%.1f, %.1f", camera.Position.X, camera.Position.Z)
	rl.DrawText(s, 10, 10, 22, rl.White)

	// Collision
	if isHittingWall() {
		camera.Position = oldCameraPosition
	}
}

func isHittingWall() bool {
	playerPos := rl.Vector2{X: camera.Position.X, Y: camera.Position.Z}
	for _, w := range walls {
		if rl.CheckCollisionCircleRec(playerPos, 0.1, w) {
			return true
		}
	}

	return false
}

func drawNeedle(start, end rl.Vector3, col color.RGBA) {
	cx := start.X - end.X
	cy := start.Z - end.Z
	cx, cy = normalize(cx, cy, compassSize)
	rl.DrawLine(compassX, compassY, int32(compassX+cx), int32(compassY+cy), col)
}

func normalize(cx, cy, size float32) (float32, float32) {
	l := float32(math.Sqrt(float64(cx*cx + cy*cy)))
	return cx / l * size, cy / l * size
}

func coord(v int32) float32 {
	return float32(v) * 4
}

func drawWall(x, y, z, w, h, t float32, col color.RGBA) {
	walls = append(walls, rl.Rectangle{X: x, Y: z, Width: w, Height: t})
	rl.DrawCube(rl.NewVector3(x, y, z), w, h, t, col)
	rl.DrawCube(rl.NewVector3(x, 3, z), w, 1, t, rl.Red) // Draw roof
}
