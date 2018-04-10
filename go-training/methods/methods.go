package methods

import (
	"strconv"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return "point: x=" + strconv.Itoa(p.X) + ", y=" + strconv.Itoa(p.Y)
}