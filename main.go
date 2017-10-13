package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	file, err := os.Create("icon.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	image := image.NewRGBA(image.Rect(0, 0, 16, 16))
	err = png.Encode(file, image)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
