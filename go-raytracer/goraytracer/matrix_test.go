package goraytracer

import (
	"testing"
)

func TestMatrixEquality(t *testing.T) {
	a := Mat4x4{
        [4][4]float64{
            {0, 1, 2, 3},
            {0, 1, 2, 3},
            {0, 1, 2.111, 3},
            {0.111, 1, 2.11, 3},
        },
    }

    b := Mat4x4{
        [4][4]float64{
            {0, 1, 2, 3},
            {0, 1, 2, 3},
            {0, 1, 2.111, 3},
            {0.111, 1, 2.11, 3},
        },
    }

    if !a.Equals(&b) {
        t.Error("Matrix equality got wrong")
    }
}

func TestMatrixMultiplication(t *testing.T) {
	a := Mat4x4{
        [4][4]float64{
            {1, 2, 3, 4},
            {5, 6, 7, 8},
            {9, 8, 7, 6},
            {5, 4, 3, 2},
        },
    }

    b := Mat4x4{
        [4][4]float64{
            {-2, 1, 2, 3},
            {3, 2, 1, -1},
            {4, 3, 6, 5},
            {1, 2, 7, 8},
        },
    }
    
    c := Mat4x4{
        [4][4]float64{
            {20, 22, 50, 48},
            {44, 54, 114, 108},
            {40, 58, 110, 102},
            {16, 26, 46, 42},
        },
    }

    mul := a.Mul(&b)

    if !mul.Equals(&c) {
        t.Error("Matrix multiplication got wrong")
    }
}

func TestIdentity(t *testing.T) {
	a := Mat4x4{
        [4][4]float64{
            {1, 2, 3, 4},
            {5, 6, 7, 8},
            {9, 8, 7, 6},
            {5, 4, 3, 2},
        },
    }

    b := NewIdentity()

    mul1 := a.Mul(b)
    mul2 := b.Mul(&a)

    if !mul1.Equals(&a) && !mul2.Equals(&a) {
        t.Error("Identity got wrong")
    }
}

func TestTranspose(t *testing.T) {
    a := Mat4x4{
        [4][4]float64{
            {0, 9, 3, 0},
            {9, 8, 0, 8},
            {1, 8, 5, 3},
            {0, 0, 5, 8},
        },
    }

    b := Mat4x4{
        [4][4]float64{
            {0, 9, 1, 0},
            {9, 8, 8, 0},
            {3, 0, 5, 5},
            {0, 8, 3, 8},
        },
    }

    c := a.T()

    if !c.Equals(&b) {
        t.Error("Transpose got wrong")
    }
}

func TestMatrixInv(t *testing.T) {
    a := Mat4x4{
        [4][4]float64{
            {0, 9, 3, 0},
            {9, 8, 0, 8},
            {1, 8, 5, 3},
            {0, 0, 5, 8},
        },
    }

    if !a.Mul(a.Inv()).Equals(NewIdentity()) || !a.Inv().Mul(&a).Equals(NewIdentity()) {
        t.Error("Inverse got wrong")
    }
}