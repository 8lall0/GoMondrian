package gomondrian

import (
	"math/rand"
	"time"
)

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func randGen() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func randInt(max, min int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// Generates a "truly" random bool value.
func (b *boolgen) randBool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}
