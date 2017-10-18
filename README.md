# Go Perfy - simple benchmarks for Go idioms

This repository contains various benchmarks that compares the performance
of various idioms in Go.

When is using a slice better than using a `map` for sets and maps?
What is the best way to map a slice of bytes to some value?
Is it more efficient to do X or is it better to do Y?

To run any benchmark, simply `cd` into the benchmark directory and run
```bash
go test -bench=. -count <# of times> > <results file>
```
You can then use [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat)
to summarize your results over `<# of times>` trials.
