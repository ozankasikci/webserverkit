package rand

import "math/rand"

// min: inclusive, max: exclusive
func RandomNumberRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func Percent(p int) bool {
	val := RandomNumberRange(0, 100)
	return p >= val
}

func PickRandomInt(values []int) int {
	i := RandomNumberRange(0, len(values))
	return values[i]
}

func PickRandomString(values []string) string {
	i := RandomNumberRange(0, len(values))
	return values[i]
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}