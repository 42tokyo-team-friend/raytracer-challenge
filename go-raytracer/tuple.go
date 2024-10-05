package goraytracer

type Tuple struct {
	x	float32
	y	float32
	z	float32
	w	float32
}

func Point(x float32, y float32, z float32) Tuple {
	return Tuple{x, y, z, 1.0};
}

func Vector(x float32, y float32, z float32) Tuple {
	return Tuple{x, y, z, 0.0};
}