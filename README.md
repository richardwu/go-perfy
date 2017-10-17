# Go Perfy - Simple Benchmarks for Go Idioms

This repository contains various benchmarks that compares the performance
of various idioms in Go.

When is a slice better than using a map for simple set membership
(and not deletions)? What is the best way to map a slice of bytes to some value?

To run any benchmark, simply `cd` into the benchmark directory and run
```
go test -bench=. -count <# of times> > <results file>
```
You can then use [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat)
to summarize your results over `<# of times>` trials.
