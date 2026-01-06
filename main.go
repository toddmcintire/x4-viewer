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
				
				if bytesRead > 0 {
					fmt.Printf("%v ", bytesRead)
					for _, v := range buf{
						fmt.Printf("%v ", v)
					}
					fmt.Println()
				}
			}

			rl.UnloadDroppedFiles()
		}

		//Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if len(filePaths) == 0 {
			rl.DrawText("Drop file", 100, 40, 20, rl.DarkGray)
		} else {
			rl.DrawTexture(texture, 0, 0, rl.RayWhite)
			//rl.DrawText("Dropped files:", 100, 40, 20, rl.DarkGray)

			//for i:=int32(0); i < int32(len(filePaths)); i++ {
				// if i%2 == 0 {
				// 	rl.DrawRectangle(0, 85+40*i, 480, 40, rl.Fade(rl.LightGray, 0.5))
				// } else {
				// 	rl.DrawRectangle(0, 85+40*i, 480, 40, rl.Fade(rl.LightGray, 0.3))
				// }
				
				// rl.DrawText(filePaths[i], 120, 100 + 40 * i, 10, rl.Gray)
				//fmt.Println(filePaths[i])
				
				//if strings.Contains(filePaths[i], "xtg") {
					//x4.GetXTGData(filePaths[i])
					//expanded := pbm.ExpandBitmap(pictureData)
					//fmt.Println(pictureData)
					// img := rl.NewImage(expanded, 480, 800, 1, rl.UncompressedGrayscale)
					// texture := rl.LoadTextureFromImage(img)
					// rl.DrawTexture(texture,0,0,rl.RayWhite)
				//}

			//}
			//rl.DrawText("drop new files...", 100, 110 + 40 * int32(filepathCounter), 20, rl.DarkGray)
		}
		rl.EndDrawing()
	}
}