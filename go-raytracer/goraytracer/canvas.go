package goraytracer

import (
	"image"
	"image/png"
	"io"
)

type Canvas struct {
	data [][]Tuple
	width int
	height int
}

func NewCanvas(width int, height int) Canvas {
	canvas := Canvas{
		make([][]Tuple, height),
		width,
		height,
	}
	for i := 0; i < height; i++ {
		canvas.data[i] = make([]Tuple, width)
	}
	return canvas
}

func (c *Canvas) Get(x int, y int) Tuple {
	return c.data[y][x]
}

func (c *Canvas) Write(x int, y int, pixel Tuple) {
	c.data[y][x] = pixel
}

func (c *Canvas) ToImage(out io.Writer) {
	rect := image.Rect(0, 0, c.width, c.height)
	img := image.NewRGBA(rect)
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			img.SetRGBA(x, y, c.Get(x, y).ToRGBA())
		}
	}
	png.Encode(out, img)
}