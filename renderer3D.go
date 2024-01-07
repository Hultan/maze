package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

	//rl.DrawCube(camera.Target, 0.5, 0.5, 0.5, rl.Purple)
	//rl.DrawCubeWires(camera.Target, 0.5, 0.5, 0.5, rl.DarkPurple)

	rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.Beige) // Draw ground
	rl.DrawCube(rl.NewVector3(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)            // Draw a blue wall
	rl.DrawCube(rl.NewVector3(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)             // Draw a green wall
	rl.DrawCube(rl.NewVector3(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)             // Draw a yellow wall

	rl.EndMode3D()
	rl.EndDrawing()
}
