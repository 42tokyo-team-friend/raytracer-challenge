package goraytracer

import (
	"image/color"
	"math"
)

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

const Epsilon = 10e-5

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (left Tuple) Add(right Tuple) Tuple {
	return Tuple{left.x + right.x, left.y + right.y, left.z + right.z, left.w + right.w}
}

func (left Tuple) Sub(right Tuple) Tuple {
	return Tuple{left.x - right.x, left.y - right.y, left.z - right.z, left.w - right.w}
}

func (t Tuple) Neg() Tuple {
	return Tuple{-t.x, -t.y, -t.z, -t.w}
}

func (t Tuple) Mul(k float64) Tuple {
	return Tuple{t.x * k, t.y * k, t.z * k, t.w * k}
}

func (t Tuple) MulTup(other Tuple) Tuple {
	return Tuple{t.x * other.x, t.y * other.y, t.z * other.z, t.w * other.w}
}

func (t Tuple) Div(k float64) Tuple {
	return Tuple{t.x / k, t.y / k, t.z / k, t.w / k}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w)
}

func (left Tuple) Dot(right Tuple) float64 {
	return left.x*right.x + left.y*right.y + left.z*right.z + left.w*right.w
}

func (left Tuple) Cross(right Tuple) Tuple {
	return Vector(left.y*right.z-left.z*right.y, left.z*right.x-left.x*right.z, left.x*right.y-left.y*right.x)
}

func (t Tuple) Normalize() Tuple {
	return t.Div(t.Magnitude())
}

func (left Tuple) Equals(right Tuple) bool {
	return math.Abs(left.x-right.x) < Epsilon &&
		math.Abs(left.y-right.y) < Epsilon &&
		math.Abs(left.z-right.z) < Epsilon &&
		math.Abs(left.w-right.w) < Epsilon
}

func (t Tuple) ToRGBA() color.RGBA {
	return color.RGBA{uint8(t.x), uint8(t.y), uint8(t.z), 255}
}

func Dot(t1, t2 Tuple) float64 {
	return t1.Dot(t2)
}

func (in Tuple) Reflect(normal Tuple) Tuple {
	return in.Sub(normal.Mul(2.0 * Dot(in, normal)))
}

func (t Tuple) Crop(k float64) Tuple {
	return Tuple{math.Min(t.x, k), math.Min(t.y, k), math.Min(t.z, k), math.Min(t.w, k)}
}
