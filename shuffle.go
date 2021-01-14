package chi_square

import (
	"math/rand"
	"time"
)

func Shuffle(s []interface{}) []interface{} {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
