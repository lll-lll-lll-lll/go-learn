package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

var space = func() string {
	result := ""
	for i := 0; i < 20; i++ {
		result += " "
	}
	return result
}()

func main() {
	picPath := "./image/background.png"
	savePath := "./shape_frame_text.png"
	err := addPictureFrameWithText(picPath, savePath,
		fmt.Sprintf("flksjeif9%sfvaskenfiew0", space))
	if err != nil {
		panic(err)
	}
}
