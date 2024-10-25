package goraytracer

import (
	"math"

	"github.com/google/uuid"
)

type Sphere struct {
	id uuid.UUID
}

func NewSphere() *Sphere {
	id, _ := uuid.NewUUID()

	return &Sphere{
		id,
	}
}

func (s *Sphere) Intersect(r *Ray) []Intersection {
	sphere_to_ray := r.origin.Sub(Point(0, 0, 0))
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(sphere_to_ray)
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