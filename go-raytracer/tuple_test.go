package goraytracer

import (
	"math"
	"testing"
)


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

func TestTupleAddition(t *testing.T) {
	t1 := Point(3, -2, 5)
	t2 := Vector(-2, 3, 1)
	ans := Tuple{1, 1, 6, 1}

	add := t1.Add(t2)

	if add != ans {
		t.Error("Addition got wrong")
	}
}

func TestTupleSubtraction(t *testing.T) {
	t1 := Point(3, 2, 1)
	t2 := Vector(5, 6, 7)
	ans := Tuple{-2, -4, -6, 1}

	sub := t1.Sub(t2)

	if sub != ans {
		t.Error("Subtraction got wrong")
	}
}

func TestTupleNegation(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	ans := Tuple{-1, 2, -3, 4}

	if a.Neg() != ans {
		t.Error("Negation got wrong")
	}
}

func TestTupleMultiplication(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	ans := Tuple{3.5, -7, 10.5, -14}

	if a.Mul(3.5) != ans {
		t.Error("Multiplicaion got wrong")
	}
}

func TestTupleDivision(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	ans := Tuple{0.5, -1, 1.5, -2}

	if a.Div(2) != ans {
		t.Error("Multiplicaion got wrong")
	}
}

func TestTupleMagnitude(t *testing.T) {
	v := Vector(-1, -2, -3)
	ans := math.Sqrt(14)

	if v.Magnitude() != ans {
		t.Error("Magnitude got wrong")
	}
}