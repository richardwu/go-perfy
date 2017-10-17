# Appending vs Indexing for Creating a Slice

## Benchmarks

There are two idioms in Go for creating a new slice of pre-determined size.

### Appending

One can preallocate the __capacity__ and `append` to it
```go
slice := make([]foo, 0, sz)
for _, e := range input {
  slice = append(slice, e)
}
```

### Indexing

One can also preallocate the __size__ and __capacity__ (given the slice's
elements are primitive enough for redundant copy assignment)
```go
slice = make([]foo, sz)
for i, e := range input {
  slice[i] = e
}
```

## Flags

```
  --slice-size      The size of the input and final slice i.e. number of
                    elements in total.
```

## Results

With `--slice-size=32` and 10 trials
```
name      time/op
Append-8  85.3ns ± 3%
Idx-8     71.6ns ± 1%
```
Indexing into the slice is more efficient. This could be attributed to
additional runtime checks of the slice's capacity per each `append` operation.
