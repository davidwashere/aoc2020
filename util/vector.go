package util

// Vector holds direction and magnitude
type Vector struct {
	X int
	Y int
	M int
}

// NewVector creates a vector
func NewVector(x, y, magnitude int) Vector {
	return Vector{
		X: x,
		Y: y,
		M: magnitude,
	}
}

// NewNormalizedVector creates a vector with magnitude 1
func NewNormalizedVector(x, y int) Vector {
	return Vector{
		X: x,
		Y: y,
		M: 1,
	}
}

// Apply will apply this vector transform to the given coords
func (v *Vector) Apply(x, y int) (tx int, ty int) {
	tx = (v.X * v.M) + x
	ty = (v.Y * v.M) + y

	return tx, ty
}
