package main

import (
	"fmt"
	"image/png"
	"os"

	dc "github.com/NineStems/dishonoredCrypt"
)

func main() {
	f, err := os.Create("outimage.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	img, err := dc.GetImageByWords("Тест")

	if err != nil {
		fmt.Println(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		fmt.Println(err)
	}
}
