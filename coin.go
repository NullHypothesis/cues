package cues

import (
	"math/rand"
)

const (
	// We arbitrarily set heads to 1 but it might as well be 0.
	heads     = 1
	maxOffset = 63
)

type coin struct {
	randBits uint64
	offset   uint
}

func newCoin() *coin {
	return &coin{}
}

// allHeads returns `true` if the next `numTosses` coin tosses are all heads.
// For 1, the probability is 0.5; for 2, it is 0.25; for 3, it is 0.125, etc.
func (r *coin) allHeads(numTosses int) bool {
	for i := 0; i < numTosses; i++ {
		// (Re)obtain randomness.
		if r.offset == 0 {
			r.randBits = rand.Uint64()
		}
		// Check if the least significant bit is heads.
		if (r.randBits>>r.offset)&1 != heads {
			r.incOffset()
			return false
		}
		r.incOffset()
	}
	return true
}

func (r *coin) incOffset() {
	r.offset = (r.offset + 1) % (maxOffset + 1)
}
