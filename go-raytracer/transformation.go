package goraytracer

import "math"

func Translation(x float64, y float64, z float64) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{1, 0, 0, x},
			{0, 1, 0, y},
			{0, 0, 1, z},
			{0, 0, 0, 1},
		},
	}
}

func Scaling(x float64, y float64, z float64) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{x, 0, 0, 0},
			{0, y, 0, 0},
			{0, 0, z, 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationX(rad float64) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{1, 0, 0, 0},
			{0, math.Cos(rad), -math.Sin(rad), 0},
			{0, math.Sin(rad), math.Cos(rad), 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationY(rad float64) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{math.Cos(rad), 0, math.Sin(rad), 0},
			{0, 1, 0, 0},
			{-math.Sin(rad), 0, math.Cos(rad), 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationZ(rad float64) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{math.Cos(rad), -math.Sin(rad), 0, 0},
			{math.Sin(rad), math.Cos(rad), 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func Shearing(xy, xz, yx, yz, zx, zy float64) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{1, xy, xz, 0},
			{yx, 1, yz, 0},
			{zx, zy, 1, 0},
			{0, 0, 0, 1},
		},
	}
}