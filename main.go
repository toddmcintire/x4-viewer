package main

import (
	"fmt"
	//"strings"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/toddmcintire/x4-viewer.git/pbm"
	"github.com/toddmcintire/x4-viewer.git/x4"
)

func main() {

	var filePaths []string
	var texture rl.Texture2D

	rl.InitWindow(480, 800, "x4 viewer")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		//Update
		if rl.IsFileDropped() {
			droppedFiles := rl.LoadDroppedFiles()
			for _, v := range droppedFiles {
				filePaths = append(filePaths, v)				
			}

			if len(filePaths) > 0 {
				buf := make([]byte, 48000)
				bytesRead := x4.GetXTGData(filePaths[0], buf)
				expanded := pbm.ExpandBitmap(buf)
				img := rl.NewImage(expanded, 480, 800, 1, rl.UncompressedGrayscale)
				texture = rl.LoadTextureFromImage(img)
			}

			rl.UnloadDroppedFiles()
		}

		//Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if len(filePaths) == 0 {
			rl.DrawText("Drop file", 200, 400, 20, rl.DarkGray)
		} else {
			rl.DrawTexture(texture, 0, 0, rl.RayWhite)
		}
		rl.EndDrawing()
	}
}