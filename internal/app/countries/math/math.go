package math

// Max returns the larger of x or y.
func Max(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}
