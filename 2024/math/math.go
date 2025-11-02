package math

func Abs(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}
