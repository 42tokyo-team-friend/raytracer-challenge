package goraytracer

import "math"

type Tuple struct {
	x	float64
	y	float64
	z	float64
	w	float64
}

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0};
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0};
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

func (t Tuple) Div(k float64) Tuple {
	return Tuple{t.x / k, t.y / k, t.z / k, t.w / k}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.x * t.x + t.y * t.y + t.z * t.z + t.w * t.w)
}