package goraytracer

type Ray struct {
	origin Tuple
	direction Tuple
}

func (r *Ray) Position(t float64) Tuple {
	return r.origin.Add(r.direction.Mul(t))
}