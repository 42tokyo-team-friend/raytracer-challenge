package goraytracer

 import "math"

type Material struct {
	Color Tuple
	Ambient float64
	Diffuse float64
	Specular float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{
		Vector(1, 1, 1),
		0.1,
		0.9,
		0.9,
		200.0,
	}
}

func Lighting(material *Material, light *PointLight, point Tuple, eyev Tuple, normalv Tuple) Tuple {
	var Diffuse Tuple
	var Ambient Tuple
	var Specular Tuple
	
	effective_color := material.Color.MulTup(light.intensity)
	lightv := light.position.Sub(point).Normalize()
	Ambient = effective_color.Mul(material.Ambient)
	light_dot_normal := Dot(lightv, normalv)

	if light_dot_normal < 0 {
		Diffuse = Vector(0, 0, 0)
		Specular = Vector(0, 0, 0)
	} else {
		Diffuse = effective_color.Mul(material.Diffuse * light_dot_normal)
		reflectv := lightv.Neg().Reflect(normalv)
		reflect_dot_eye := Dot(reflectv, eyev)
		if reflect_dot_eye <= 0 {
			Specular = Vector(1, 1, 1)
		} else {
			factor := math.Pow(reflect_dot_eye, material.Shininess)
			Specular = light.intensity.Mul(material.Specular * factor)
		}
	}
	return Ambient.Add(Diffuse).Add(Specular)
}