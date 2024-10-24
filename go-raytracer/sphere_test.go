package goraytracer

import "testing"

func TestSphereIntersectAtTwoPoint(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := NewSphere()
	xs := s.Intersect(&r)
	
	if len(xs) != 2 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}

	if xs[0] != 4.0 || xs[1] != 6.0 {
		t.Error("The ray intersects with a sphere in a wrong point")
	} 
}