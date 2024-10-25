package goraytracer

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func NewRay(origin Tuple, direction Tuple) *Ray {
	return &Ray{
		origin,
		direction,
	}
}

func (r *Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r *Ray) Transform(m *Mat4x4) *Ray {
	return &Ray{
		m.MulTup(r.Origin),
		m.MulTup(r.Direction),
	}
}
