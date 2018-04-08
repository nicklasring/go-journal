package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func screenShot(txtFileName string) (fileName string) {
	bounds := image.Rectangle{}
	bounds.Min, bounds.Max = getMouseCoords()

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	fileName = fmt.Sprintf("%s.png", txtFileName[:len(txtFileName)-4])
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	return fileName
}
