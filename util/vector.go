package util

// Vector holds direction and magnitude
type Vector struct {
	X int
	Y int
	M int
}

// Apply will apply this vector transform to the given coords
func (v *Vector) Apply(x, y int) (tx int, ty int) {
	tx = (v.X * v.M) + x
	ty = (v.Y * v.M) + y

	return tx, ty
}
