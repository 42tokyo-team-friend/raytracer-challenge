package goraytracer

import (
	"os"
	"testing"
)

func TestCanvasCreation(t *testing.T) {
	canvas := NewCanvas(10, 20)

	if canvas.width != 10 {
		t.Error("width is not correctly set")
	}

	if canvas.height != 20 {
		t.Error("height is not correctly set")
	}

	for i := 0; i < canvas.height; i++ {
		for j := 0; j < canvas.width; j++ {
			if !canvas.data[i][j].Equals(Vector(0, 0, 0)){
				t.Error("All pixels must be initialized as zero")
			}
		}
	}
}


func TestCanvasGet(t *testing.T) {
	c := NewCanvas(10, 20)
	red := Vector(1, 0, 0)
	c.data[3][2] = red;

	if !c.Get(2, 3).Equals(red) {
		t.Error("canvas.Get got wrong")
	}
}


func TestCanvasWrite(t *testing.T) {
	c := NewCanvas(10, 20)
	red := Vector(1, 0, 0)
	c.Write(2, 3, red)

	if !c.Get(2, 3).Equals(red) {
		t.Error("canvas.Write got wrong")
	}
}

func TestCanvasToImage(t *testing.T) {
	c := NewCanvas(5, 3)
	f, err := os.Create("TestCanvasToImage.jpeg")

	if err != nil {
		t.Error("Open failed")
	}

	c.ToImage(f)

	f.Close();
}