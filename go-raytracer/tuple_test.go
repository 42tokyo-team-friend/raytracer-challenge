package goraytracer

import "testing"


func TestPointCreation(t *testing.T) {
	point := Point(4, -4, 3);

	if point.w != 1.0 {
		t.Error("W coordinate of point must be 1.0");
	}
}

func TestVectorCreation(t *testing.T) {
	vector := Vector(4, -4, 3);

	if vector.w != 0.0 {
		t.Error("W coordinate of vector must be 0.0");
	}
}