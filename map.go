package c80

import (
	"image"
)

func Map(width int, height int, x int, y int) image.Image {
	return machine.Map(width, height, x, y)
}

func DrawMap(screenX int, screenY int, screenWidth int, screenHeight int, mapX int, mapY int) {
	machine.DrawMap(screenX, screenY, screenWidth, screenHeight, mapX, mapY)
}
