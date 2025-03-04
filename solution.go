package square

import "math"

type intCustomType int

const (
	SidesCircle intCustomType = iota
	_
	_
	SidesTriangle
	SidesSquare
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var s float64
	switch sidesNum {
	case SidesCircle:
		s = math.Pi * sideLen * sideLen
	case SidesTriangle:
		s = sideLen * sideLen * math.Sqrt(3) / 4
	case SidesSquare:
		s = sideLen * sideLen
	}
	return s
}
