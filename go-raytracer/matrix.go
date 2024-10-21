package goraytracer

import "math"

type Mat4x4 struct {
	data [4][4]float64
}

func (left *Mat4x4) Equals(right *Mat4x4) bool {
	flag := true

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if math.Abs(left.data[i][j] - right.data[i][j]) > Epsilon {
				flag = false
			}
		}
	}

	return flag
}

func (left *Mat4x4) Mul(right *Mat4x4) Mat4x4 {
	return Mat4x4{
		[4][4]float64{
			{
				left.data[0][0] * right.data[0][0] + left.data[0][1] * right.data[1][0] + left.data[0][2] * right.data[2][0] + left.data[0][3] * right.data[3][0],
				left.data[0][0] * right.data[0][2] + right.data[0][1] * right.data[1][1] + left.data[0][2] * right.data[2][1] + left.data[0][3] * right.data[3][1],
				left.data[0][0] * right.data[0][2] + right.data[0][2] * right.data[1][2] + left.data[0][2] * right.data[2][2] + left.data[0][3] * right.data[3][2],
				left.data[0][0] * right.data[0][3] + right.data[0][3] * right.data[1][3] + left.data[0][2] * right.data[2][3] + left.data[0][3] * right.data[3][3],
			},
			{
				left.data[1][0] * right.data[0][0] + left.data[1][1] * right.data[1][0] + left.data[1][2] * right.data[2][0] + left.data[1][3] * right.data[3][0],
				left.data[1][0] * right.data[0][2] + left.data[1][1] * right.data[1][1] + left.data[1][2] * right.data[2][1] + left.data[1][3] * right.data[3][1],
				left.data[1][0] * right.data[0][2] + left.data[1][2] * right.data[1][2] + left.data[1][2] * right.data[2][2] + left.data[1][3] * right.data[3][2],
				left.data[1][0] * right.data[0][3] + left.data[1][3] * right.data[1][3] + left.data[1][2] * right.data[2][3] + left.data[1][3] * right.data[3][3],
			},
			{
				left.data[2][0] * right.data[0][0] + left.data[2][1] * right.data[1][0] + left.data[2][2] * right.data[2][0] + left.data[2][3] * right.data[3][0],
				left.data[2][0] * right.data[0][2] + left.data[2][1] * right.data[1][1] + left.data[2][2] * right.data[2][1] + left.data[2][3] * right.data[3][1],
				left.data[2][0] * right.data[0][2] + left.data[2][2] * right.data[1][2] + left.data[2][2] * right.data[2][2] + left.data[2][3] * right.data[3][2],
				left.data[2][0] * right.data[0][3] + left.data[2][3] * right.data[1][3] + left.data[2][2] * right.data[2][3] + left.data[2][3] * right.data[3][3],
			},
			{
				left.data[3][0] * right.data[0][0] + left.data[3][1] * right.data[1][0] + left.data[3][2] * right.data[2][0] + left.data[3][3] * right.data[3][0],
				left.data[3][0] * right.data[0][2] + left.data[3][1] * right.data[1][1] + left.data[3][2] * right.data[2][1] + left.data[3][3] * right.data[3][1],
				left.data[3][0] * right.data[0][2] + left.data[3][2] * right.data[1][2] + left.data[3][2] * right.data[2][2] + left.data[3][3] * right.data[3][2],
				left.data[3][0] * right.data[0][3] + left.data[3][3] * right.data[1][3] + left.data[3][2] * right.data[2][3] + left.data[3][3] * right.data[3][3],
			},
		},
	}
}