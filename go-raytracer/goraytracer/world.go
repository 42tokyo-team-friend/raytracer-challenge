package goraytracer

type World struct {
	Objects []*Sphere
	Lights []*PointLight
}

func NewWorld(Objects []*Sphere, Lights []*PointLight) *World {
	return &World{
		Objects,
		Lights,
	}
}

func DefaultWorld() *World {
	s1 := NewSphere()
	s1m := NewMaterial()
	s1m.Color = Vector(0.8, 1.0, 0.6)
	s1m.Diffuse = 0.7
	s1m.Specular = 0.2
	s1.SetMaterial(s1m)
	s2 := NewSphere()
	s2.SetTransform(Scaling(0.5, 0.5, 0.5))
	light := NewPointLight(Point(-10, 10, -10), Vector(1, 1, 1))
	return NewWorld([]*Sphere{s1}, []*PointLight{light})
}
