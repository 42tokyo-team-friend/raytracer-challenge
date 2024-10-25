package main

import (
	rt "goraytracer"
	"os"
)

func main() {
	origin := rt.Point(0, 0, -5)
	wallZ := 10.0
	wallSizes := 7.0
	canvasPixels := 1000
	pixelSize := wallSizes / float64(canvasPixels)
	half := wallSizes / 2.0
	s := rt.NewSphere()
	c := rt.NewCanvas(canvasPixels, canvasPixels)
	m := rt.NewMaterial()
	m.Color = rt.Vector(1, 0.2, 1)
	s.SetMaterial(m)
	light := rt.NewPointLight(rt.Point(-10, 10, -10), rt.Vector(1, 1, 1))

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			p := rt.Point(worldX, worldY, wallZ)
			r := rt.NewRay(origin, p.Sub(origin).Normalize())
			xs := s.Intersect(r)
			hit := rt.Hit(xs)

			if !hit.Empty() {
				point := r.Position(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := r.Direction.Neg()
				color := rt.Lighting(hit.Object.Material, light, point, eye, normal).Mul(255).Crop(255)
				c.Write(x, y, color)
			} else {
				c.Write(x, y, rt.Vector(0, 0, 0))
			}
		}
	}

	f, _ := os.Create("TestCanvasToImage.png")

	c.ToImage(f)

	f.Close()
}
