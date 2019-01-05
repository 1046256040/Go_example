package mathUtil

import (
	"math/rand"
	"time"
)

func RandInt(a, b int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(a)
	return n
}
