package goraytracer

import "testing"

func TestRayPosition(t *testing.T) {
	r := Ray{origin: Point(2, 3, 4), direction: Vector(1, 0, 0)}

	if !r.Position(2.5).Equals(Point(4.5, 3, 4)) {
		t.Error("Ray position got wrong")
	}
}

func TestTranslatingRay(t *testing.T) {
	r := Ray{Point(1, 2, 3), Vector(0, 1, 0)}
	m := Translation(3, 4, 5)
	r2 := r.Transform(m)

	if !r2.origin.Equals(Point(4, 6, 8)) {
		t.Error("origin is not correctly transformed.")
	}

	if !r2.direction.Equals(Vector(0, 1, 0)) {
		t.Error("direction is not correctly transformed.")
	}
}

func TestScalingRay(t *testing.T) {
	r := Ray{Point(1, 2, 3), Vector(0, 1, 0)}
	m := Scaling(2, 3, 4)
	r2 := r.Transform(m)

	if !r2.origin.Equals(Point(2, 6, 12)) {
		t.Error("origin is not correctly transformed.")
	}

	if !r2.direction.Equals(Vector(0, 3, 0)) {
		t.Error("direction is not correctly transformed.")
	}
}