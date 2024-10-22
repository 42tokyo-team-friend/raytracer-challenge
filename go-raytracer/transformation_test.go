package goraytracer

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	if !Translation(5, -3, 2).MulTup(Point(-3, 4, 5)).Equals(Point(2, 1, 7)) {
		t.Error("Translation got wrong")
	}

	if !Translation(5, -3, 2).Inv().MulTup(Point(-3, 4, 5)).Equals(Point(-8, 7, 3)) {
		t.Error("Translation got wrong")
	}
}

func TestTranslationVector(t *testing.T) {
	if !Translation(5, -3, 2).MulTup(Vector(-3, 4, 5)).Equals(Vector(-3, 4, 5)) {
		t.Error("Translation does not affect vectors.")
	}
}

func TestScaling(t *testing.T) {
	if !Scaling(2, 3, 4).MulTup(Point(-4, 6, 8)).Equals(Point(-8, 18, 32)) {
		t.Error("Scaling applied to Point got wrong")
	}

	if !Scaling(2, 3, 4).MulTup(Vector(-4, 6, 8)).Equals(Vector(-8, 18, 32)) {
		t.Error("Scaling applied to Vector got wrong")
	}

	if !Scaling(2, 3, 4).Inv().MulTup(Vector(-4, 6, 8)).Equals(Vector(-2, 2, 2)) {
		t.Error("Inverse Scaling got wrong")
	}
}

func TestRotationX(t *testing.T) {
	if !RotationX(math.Pi / 4).MulTup(Point(0, 1, 0)).Equals(Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)) {
		t.Error("Rotation around x axis got wrong")
	}

	if !RotationX(math.Pi / 4).Inv().MulTup(Point(0, 1, 0)).Equals(Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)) {
		t.Error("The inverse of x-rotation rotates in opposite direction.")
	}
}

func TestRotationY(t *testing.T) {
	if !RotationY(math.Pi / 4).MulTup(Point(0, 0, 1)).Equals(Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)) {
		t.Error("Rotation around x axis got wrong")
	}
}

func TestRotationZ(t *testing.T) {
	if !RotationZ(math.Pi / 4).MulTup(Point(0, 1, 0)).Equals(Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)) {
		t.Error("Rotation around y axis got wrong")
	}
}