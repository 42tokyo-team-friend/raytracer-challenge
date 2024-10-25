package goraytracer

import (
	"testing"
	"math"
)

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

func TestIntersectingScaledSphere(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := NewSphere()
	s.transform = Scaling(2, 2, 2)
	xs := s.Intersect(&r)

	if len(xs) != 2 {
		t.Error("The ray intersects with a sphere in a wrog number of points")
	}

	if xs[0].t != 3.0 || xs[1].t != 7.0 {
		t.Error("The ray intersects with a sphere in a wrong point")
	} 
}

func TestNormalXAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(1, 0, 0))
	
	if !n.Equals(Vector(1, 0, 0)) {
		t.Error("Normal is computed wrongly.")
	}
}

func TestNromalAtNonaxialPoint(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(Point(math.Sqrt(3) / 3, math.Sqrt(3) / 3, math.Sqrt(3) / 3))

	if !n.Equals(Vector(math.Sqrt(3) / 3, math.Sqrt(3) / 3, math.Sqrt(3) / 3).Normalize()) {
		t.Error("Normal is computed wrongly.")
	}
}

func TestNormalOnTranslated(t *testing.T) {
	s := NewSphere()
	s.SetTransform(Translation(0, 1, 0))
	n := s.NormalAt(Point(0, 1.70711, -0.70711))

	if !n.Equals(Vector(0, 0.70711, -0.70711)) {
		t.Error("Normal is computed wrongly when translated.")
	}
}

func TestNormalOnTransformed(t *testing.T) {
	s := NewSphere()
	s.SetTransform(Scaling(1, 0.5, 1).Mul(RotationZ(math.Pi / 5)))
	n := s.NormalAt(Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	
	if !n.Equals(Vector(0, 0.97014, -0.24254)) {
		t.Error("Normal is computed wrongly when transformed.")
	}
}