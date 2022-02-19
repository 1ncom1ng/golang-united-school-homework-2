package square

import (
	"math"
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {

	const SidesTriangle int = 3
	const SidesSquare int = 4
	const SidesCircle int = 0
	if ( sidesNum == 0 ) { 
		square = Pi * sideLen * sideLen
		} 
		else if ( sidesNum == 3 )  { 
			square = math.Sqrt(3) * sideLen * sideLen / 4
		}
			else if ( sidesNum == 4 )  {
			square = sideLen * sideLen 
			}
				else { square = 0 }


	return square
}
