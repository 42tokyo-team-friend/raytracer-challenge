package goraytracer

type Intersection struct {
	T      float64
	Object *Sphere
}

func Hit(intersections []Intersection) Intersection {
	var it Intersection

	m := -1.0

	for i := 0; i < len(intersections); i++ {
		if intersections[i].T > 0 && m == -1 {
			m = intersections[i].T
			it = intersections[i]
		}
		if intersections[i].T > 0 && intersections[i].T < m {
			m = intersections[i].T
			it = intersections[i]
		}
	}
	return it
}

func (it *Intersection) Empty() bool {
	return it.Object == nil
}
