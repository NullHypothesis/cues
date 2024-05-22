package cues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargeSeq(t *testing.T) {
	var (
		c         = newCoin()
		numTosses = 1_000
	)

	// The odds of getting 1,000 heads in a row are negligible.
	assert.Equal(t, false, c.allHeads(numTosses))
}

func TestEventualHead(t *testing.T) {
	var (
		c         = newCoin()
		numTosses = 1_000
	)

	// The odds of never getting heads after 1,000 tosses are equally
	// negligible.
	for i := 0; i < numTosses; i++ {
		if isHead := c.allHeads(1); isHead {
			return
		}
	}
	t.Errorf("No heads after %d tosses.", numTosses)
}

func TestReseed(t *testing.T) {
	var c = newCoin()

	_ = c.allHeads(1)
	origState := c.randBits

	for i := 1; i <= maxOffset; i++ {
		_ = c.allHeads(1)
	}
	// At this point, we're at the last bit of the original state.
	assert.Equal(t, origState, c.randBits)

	// After one more toss, the state should have changed.
	_ = c.allHeads(1)
	assert.NotEqual(t, origState, c.randBits)
}

func TestOffset(t *testing.T) {
	var (
		c      = newCoin()
		i uint = 0
	)

	// At the beginning, the offset should be 0.
	assert.Equal(t, i, c.offset)

	// With each toss, we expect the offset to increment.
	for i = 0; i < maxOffset; i++ {
		assert.Equal(t, i, c.offset)
		_ = c.allHeads(1)
	}

	// With the next toss, we expect the offset to wrap back to 0.
	_ = c.allHeads(1)
	assert.Equal(t, uint(0), c.offset)
}
