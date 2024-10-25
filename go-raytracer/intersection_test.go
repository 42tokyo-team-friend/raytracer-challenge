package goraytracer

import "testing"

func TestHitAllPositiveT(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := []Intersection{i1, i2}
	i := hit(xs)

	if i.t != 1 {
		t.Error("hit got wrong")
	}
}

func TestHitWhenHaveNegative(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := []Intersection{i1, i2}
	i := hit(xs)

	if i.t != 1 {
		t.Error("hit got wrong")
	}
}

func TestHitAllNegative(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := []Intersection{i1, i2}
	i := hit(xs)

	if i.object != nil {
		t.Error("hit got wrong")
	}
}