package goraytracer

import (
	"math"

	"github.com/google/uuid"
)

type Sphere struct {
	id uuid.UUID
	transform *Mat4x4
}

func NewSphere() *Sphere {
	id, _ := uuid.NewUUID()

	return &Sphere{
		id,
		NewIdentity(),
	}
}

func (s *Sphere) SetTransform(m *Mat4x4) {
	s.transform = m
}

func (s *Sphere) Intersect(r *Ray) []Intersection {
	r2 := r.Transform(s.transform.Inv())
	sphere_to_ray := r2.origin.Sub(Point(0, 0, 0))
	a := r2.direction.Dot(r2.direction)
	b := 2 * r2.direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1

	D := b * b - 4 * a * c
	
	if D < 0 {
		return []Intersection{}
	}

	return []Intersection{
		{(-b - math.Sqrt(D)) / (2 * a), s},
		{(-b + math.Sqrt(D)) / (2 * a), s},
	}
}