package goraytracer

type Intersection struct {
	t float64
	object *Sphere
}

func Hit(intersections []Intersection) Intersection {
	var it Intersection

	min := -1.0

	for i := 0; i < len(intersections); i++ {
		if intersections[i].t > 0 && min == -1 {
			min = intersections[i].t
			it = intersections[i]
		}
		if intersections[i].t > 0 && intersections[i].t < min {
			min = intersections[i].t
			it = intersections[i]
		}
	}
	return it
}

func (it *Intersection) Empty() bool {
	return it.object == nil
}