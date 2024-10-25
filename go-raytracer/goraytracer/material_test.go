package goraytracer

import (
	"testing"
	"math"
)

func TestLightningEyeBetweenightAndSurface(t *testing.T) {
	m := NewMaterial()
	position := Point(0, 0, 0)
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), Vector(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)

	if !result.Equals(Vector(1.9, 1.9, 1.9)) {
		t.Error("lightning got wrong.")
	}
}

func TestLightningWithEyeOppositeSurface(t *testing.T) {
	m := NewMaterial()
	position := Point(0, 0, 0)
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 10, -10), Vector(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)

	if !result.Equals(Vector(0.7364, 0.7364, 0.7364)) {
		t.Error("lightning with eye opposite surface got wrong.")
	}
}

func TestLightningEyeInPathOfReflectionVector(t *testing.T) {
	m := NewMaterial()
	position := Point(0, 0, 0)
	eyev := Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 10, -10), Vector(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)

	if !result.Equals(Vector(1.6364, 1.6364, 1.6364)) {
		t.Error("lightning with eye in the path of the reflection vector got wrong")
	}
}

func TestLightningEyeBehindSurface(t *testing.T) {
	m := NewMaterial()
	position := Point(0, 0, 0)
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, 10), Vector(1, 1, 1))
	result := Lighting(m, light, position, eyev, normalv)

	if !result.Equals(Vector(0.1, 0.1, 0.1)) {
		t.Error("lightning with eye behind the surface")
	}
}