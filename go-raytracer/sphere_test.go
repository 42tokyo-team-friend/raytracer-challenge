package goraytracer

import "testing"

func TestSphereIntersectAtTwoPoint(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := NewSphere()
	xs := s.Intersect(&r)
	
	if len(xs) != 2 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}

	if xs[0].t != 4.0 || xs[1].t != 6.0 {
		t.Error("The ray intersects with a sphere in a wrong point")
	}

	if xs[0].object != s || xs[1].object != s {
		t.Error("The ray intersects with a wrong sphere")
	}
}

func TestSphereIntersectionAtTangent(t *testing.T) {
	r := Ray{Point(0, 1, -5), Vector(0, 0, 1)}
	s := NewSphere()
	xs := s.Intersect(&r)

	if len(xs) != 2 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}

	if xs[0].t != 5.0 || xs[1].t != 5.0 {
		t.Error("The ray intersects with a sphere in a wrong point")
	} 
}


func TestMissingSphere(t *testing.T) {
	r := Ray{Point(0, 2, -5), Vector(0, 0, 1)}
	s := NewSphere()
	xs := s.Intersect(&r)

	if len(xs) != 0 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	r := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	s := NewSphere()
	xs := s.Intersect(&r)

	if len(xs) != 2 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}

	if xs[0].t != -1.0 || xs[1].t != 1.0 {
		t.Error("The ray intersects with a sphere in a wrong point")
	} 
}

func TestSphereBehindRay(t *testing.T) {
	r := Ray{Point(0, 0, 5), Vector(0, 0, 1)}
	s := NewSphere()
	xs := s.Intersect(&r)

	if len(xs) != 2 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}

	if xs[0].t != -6.0 || xs[1].t != -4.0 {
		t.Error("The ray intersects with a sphere in a wrong point")
	} 
}