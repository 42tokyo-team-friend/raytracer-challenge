package goraytracer

type PointLight struct {
	position Tuple
	intensity Tuple
}

func NewPointLight(position Tuple, intensity Tuple) *PointLight {
	return &PointLight{
		position,
		intensity,
	}
}