package cues

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFewElems(t *testing.T) {
	var (
		bufSize = 128
		c       = New[int](bufSize + 1)
	)

	// If we ingest fewer elements than our buffer can hold, we should get a
	// precise estimate.
	for i := 0; i < bufSize; i++ {
		c.Feed(i)
	}
	assert.Equal(t, uint64(bufSize), c.Estimate())
}

func TestManyElems(t *testing.T) {
	var (
		bufSize = 128
		c       = New[int](bufSize)
	)

	for i := 0; i < 1_000; i++ {
		c.Feed(i)
	}
	e := c.Estimate()
	// Make sure that the estimate is within a reasonable range.
	assert.Greater(t, e, uint64(500))
	assert.Less(t, e, uint64(1_500))
}

func BenchmarkCues(b *testing.B) {
	c := New[int](1_024)
	for i := 0; i < b.N; i++ {
		c.Feed(rand.Intn(100_000))
	}
}
