package goraytracer

type Ray struct {
	origin Tuple
	direction Tuple
}

func NewRay(origin Tuple, direction Tuple) *Ray {
	return &Ray{
		origin,
		direction,
	}
}

func (r *Ray) Position(t float64) Tuple {
	return r.origin.Add(r.direction.Mul(t))
}

func (r *Ray) Transform(m *Mat4x4) *Ray {
	return &Ray{
		m.MulTup(r.origin),
		m.MulTup(r.direction),
	}
}