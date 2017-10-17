# Pre-allocating slices

It is more efficient to pre-allocate a slice's capacity than to naively
append. This benchmark compares whether it is more efficient to pre-calculate
the number of elements that will be in the final slice from an input slice
with some predicate that evaluates to true or false.

That is, we have some `input []bool` that has `true` and `false` elements,
and we want to only append the given index `i` if `input[i] == true`.

## Benchmarks

### Naive appending

```go
var input []bool
var output []int
...
for i, e := range input {
  if e {
    output = append(output, i)
  }
}
...
```

### Pre-computing size

```go
var input []bool
...
count := 0
for i, e := range input {
  if e {
    count++
  }
}

output = make([]int, count)
idx := 0
for i, e := range input {
  if e {
    output[idx] = i
    idx++
  }
}
...
```

### Maximum allocation

```go
var input []bool
...
output := make([]int, 0, len(input))
for i, e := range input {
  if e {
    output = append(output, i)
  }
}
...
```

## Flags
```
  --max-elems       The maximum number of elements that can be in the final
                    slice (the size of the input slice).
  --append-pct      The percentage of elements (in the input slice) that will
                    be appended to the final slice.
```

## Results

With `--max-elems=32`, `--append-pct=0.9` and 10 trials
```
name          time/op
Append-8      1.19µs ± 0%
ExactAlloc-8  1.07µs ±13%
MaxAlloc-8    1.06µs ± 6%
```
Maximum allocation always wins out but one may allocate unnecessary memory if
`--append-pct` is low. There is some value to exact allocation if `--append-pct`
is high.
