package code

import (
	"math/rand"
	"time"
)

//Return random number
func Random(min, max int) (number int) {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(max-min) + min
	return
}
