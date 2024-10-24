package goraytracer

import "testing"

func TestRayPosition(t *testing.T) {
	r := Ray{origin: Point(2, 3, 4), direction: Vector(1, 0, 0)}

	if !r.Position(2.5).Equals(Point(4.5, 3, 4)) {
		t.Error("Ray position got wrong")
	}
}