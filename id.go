package animal

import (
	"math/rand"
	"strings"
	"time"
)

// NewID returns an ID of form adjetive + adjetive noun
func NewID() string {
	rand.Seed(time.Now().Unix())

	animal := Animals[rand.Intn(len(Animals))]
	adj := Adjetive1[rand.Intn(len(Adjetive1))]

	return strings.ToLower(animal + adj)
}
