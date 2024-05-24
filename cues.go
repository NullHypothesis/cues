package cues

import (
	"math"
)

type Counter[T comparable] struct {
	round   int
	bufSize int
	elems   map[T]struct{}
	coin    *coin
}

// New returns a new Counter with the given buffer size.  Larger buffer sizes
// yield more accurate estimates.
func New[T comparable](bufSize int) *Counter[T] {
	return &Counter[T]{
		elems:   make(map[T]struct{}),
		bufSize: bufSize,
		round:   0,
		coin:    newCoin(),
	}
}

// Feed adds the given `elem` to the counter.
func (c *Counter[T]) Feed(elem T) {
	c.elems[elem] = struct{}{}
	// After the initial round, keep the element with probability 0.5^round.
	if c.round > 0 && !c.coin.allHeads(c.round) {
		delete(c.elems, elem)
	}

	// Once the buffer is full, prune each element with probability 0.5.
	if len(c.elems) == c.bufSize {
		for e := range c.elems {
			if c.coin.allHeads(1) {
				delete(c.elems, e)
			}
		}
		c.round++
	}
}

// Estimate returns the estimated number of unique elements that were fed into
// the counter.
func (c *Counter[T]) Estimate() uint64 {
	p := 1 / (math.Pow(2, float64(c.round)))
	return uint64(float64(len(c.elems)) / p)
}
