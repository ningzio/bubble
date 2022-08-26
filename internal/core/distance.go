package core

import "fmt"

type Distance struct {
	distance int8
}

func NewDistance(p1, p2 *Pitch) *Distance {
	x := p1.Index() - p2.Index()
	if x < 0 {
		x *= -1
	}
	return &Distance{distance: x}
}

func (d *Distance) String() string {
	if int(d.distance) > len(intToCN) {
		return "9度以上超纲了"
	}
	return fmt.Sprintf("%s度", intToCN[d.distance])
}
