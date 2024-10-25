package main 

import (
	rt "goraytracer"
	"os"
)

func main() {
	width := 200
	height := 200
	origin := rt.Point(0, 0, -6)
	s := rt.NewSphere()
	c := rt.NewCanvas(width, height)
	m := rt.NewMaterial()
	m.Color = rt.Vector(1, 0.2, 1)
	s.SetMaterial(m)
	s.SetTransform(rt.Scaling(5.95, 5.95, 5.95))
	light := rt.NewPointLight(rt.Point(-10, 10, -10), rt.Vector(1, 1, 1))

	for y := height / 2 - 1; y > -height / 2; y-- {
		for x := width / 2 - 1; x > -width / 2; x-- {
			p := rt.Point(float64(x), float64(y), 0)
			r := rt.NewRay(origin, p.Sub(origin).Normalize())
			xs := s.Intersect(r)
			hit := rt.Hit(xs)
			point := r.Position(hit.t)
			
			if !hit.Empty() {
				c.Write(x + width / 2, y + height / 2, rt.Vector(1, 0, 0).Mul(255))
			} else {
				c.Write(x + width / 2, y + height / 2, rt.Vector(0, 0, 0).Mul(255))
			}
		}
	}
	
	f, _ := os.Create("TestCanvasToImage.jpeg")

	c.ToImage(f)

	f.Close();
}