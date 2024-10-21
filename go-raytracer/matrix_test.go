package goraytracer

import "testing"

func TestMatrixEquality(t *testing.T) {
	a := Mat4x4{
        [4][4]float64{
            {0, 1, 2, 3},
            {0, 1, 2, 3},
            {0, 1, 2.111, 3},
            {0.111, 1, 2.11, 3},
        }
    }

    b := Mat4x4{
        [4][4]float64{
            {0, 1, 2, 3},
            {0, 1, 2, 3},
            {0, 1, 2.111, 3},
            {0.111, 1, 2.11, 3},
        }
    }
}