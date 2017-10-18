# []byte vs bytes.Buffer

## Benchmarks

To create a new byte slice from an aribtrary number of ingress bytes, one
can either `append` to a byte slice (`[]byte`) or `Write` to a `bytes.Buffer`.

### Unallocated byte slice

```go
var out []byte
var foo []byte
for {
  ...
  out = append(out, foo...)
  ...
}
```

### Pre-allocated byte slice

```go
var input [][]byte
...
out := make([]byte, 0, len(input)*maxSize(input))
for _, foo := range input {
  out = append(out, foo...)
}
```
This method requires knowing the total size of the all input bytes or the upper
bound.

### Bytes buffer

```go
import "bytes"

var out bytes.Buffer
var foo []byte
for {
  ...
  out = append(out, foo...)
  ...
}
```

## Flags

```
  --num-elems       The number of `[]byte` elements in the input
  --elem-size       The length of each `[]byte` element
```

## Results

### Go 1.9

With `--num-elems=100`, `--elem-size=32` and 10 trials
```
name             time/op
UnallocSlice-8   237µs ± 4%
PreallocSlice-8  238µs ± 2%
BytesBuffer-8    236µs ± 2%
```
`bytes.Buffer` doesn't seem to offer any advantage over a raw `[]byte` slice
when simply appending and generating a `[]byte` slice from a stream of `[]byte`
slices.
