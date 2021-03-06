package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// Version is the version of this application.
var Version = ""

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Println(Version)
			os.Exit(0)
		case "--help", "-h":
			fmt.Println(help())
			os.Exit(0)
		}
	}

	file, err := os.Create("icon.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	scale := 1
	if len(os.Args) > 1 {
		scale, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	image := image.NewRGBA(image.Rect(0, 0, 16*scale, 16*scale))
	drawDots(image, scale)

	err = png.Encode(file, image)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func help() string {
	lines := []string{}
	lines = append(lines, "Usage:")
	lines = append(lines, "  icon [scale]")
	lines = append(lines, "  icon (--version | -v)")
	lines = append(lines, "  icon (--help | -h)")
	lines = append(lines, "")
	lines = append(lines, "Options:")
	lines = append(lines, "  --version, -v\tShow version number")
	lines = append(lines, "  --help, -h\tShow help message")

	return strings.Join(lines, "\n")
}

var dots = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
	{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
	{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func drawDots(image *image.RGBA, scale int) {
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	for x := 0; x < 16; x++ {
		for i := 0; i < scale; i++ {
			for y := 0; y < 16; y++ {
				for j := 0; j < scale; j++ {
					if dots[y][x] == 1 {
						image.Set(x*scale+i, y*scale+j, white)
					} else {
						image.Set(x*scale+i, y*scale+j, black)
					}
				}
			}
		}
	}
}
