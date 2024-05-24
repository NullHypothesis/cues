# cues

[![GoDoc](https://pkg.go.dev/badge/github.com/NullHypothesis/cues)](https://pkg.go.dev/github.com/NullHypothesis/cues)
[![Go Report Card](https://goreportcard.com/badge/github.com/NullHypothesis/cues)](https://goreportcard.com/report/github.com/NullHypothesis/cues)

This Go package implements
[Chakraborty,
Vinodchandran, and
Meel's
algorithm for counting distinct elements in a
stream](https://arxiv.org/pdf/2301.10191).

Assume you are dealing with a large number of elements –
larger than you could possibly store in memory.
You now want to know approximately how many unique elements you have.
This algorithm solves that problem.
The only configurable parameter is the algorithm's buffer size.
The larger the buffer,
the more accurate the algorithm's estimate.

To use this package,
first initialize a new counter:

```go
const bufSize = 1024
c := cues.New(bufSize)
```

Next, feed whatever objects you're dealing with into the counter.
Elements must be of type [`comparable`](https://go.dev/blog/comparable).

```go
for _, e := range elems {
    c.Feed(e)
}
```

Finally, obtain an estimate of the number of unique elements in your stream:

```go
fmt.Printf("Approximately %d unique elements.\n", c.Estimate())
```

For a full example,
take a look at [main.go](cmd/main.go),
which is included in this package.