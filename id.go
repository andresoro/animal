package animal

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// NewID returns an ID of form adjetive + animal noun
func New() string {

	animal := capitalize(Animals[rand.Intn(len(Animals))])
	adj := capitalize(Adjetive1[rand.Intn(len(Adjetive1))])
	adj2 := capitalize(Adjetive2[rand.Intn(len(Adjetive2))])

	return (adj + adj2 + animal)
}

func capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}
