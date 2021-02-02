package two

import (
	"math/rand"
	"time"
)

func GetFilePermission() int {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return r.Intn(8)
}