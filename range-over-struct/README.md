# Index vs copy range (iterating over a slice)

## Benchmarks

In Go, there are two possible ways to access elements in a slice: indexing
into the slice or accessing the "copied" element in the `range` loop.

### Indexing

```go
foo = make([]largeStruct, sz)
...
for i := range foo {
  doSomething(foo[i])
}
```

### Indexing

```go
foo = make([]largeStruct, sz)
...
for _, elem := range foo {
  doSomething(elem)
}
```

## Flags

```
  --num-structs     The number of largeStructs initialized in the input slice.
```

## Results

### Go 1.9

With `--num-structs=200`, `16 bytes` of data per `largeStruct`, and 10 trials
```
name                   time/op
IndexRange-8            128ns ± 3%
CopyRange-8             230ns ± 3%
IndexRangeEveryByte-8  2.24µs ± 1%
CopyRangeEveryByte-8   1.89µs ± 1%
```
The tradeoff between indexing and copying (across both cases where we range
over just the slice and where we range over the slice and internal data `byte array`)
seems to hold for larger struct sizes.

Here are the results with `256 bytes` of data per `largeStruct`
```
name                   time/op
IndexRange-8           1.91µs ± 3%
CopyRange-8            5.57µs ± 3%
IndexRangeEveryByte-8  36.5µs ± 1%
CopyRangeEveryByte-8   24.1µs ± 1%
```
with `4096 bytes` (4KB)
```
name                   time/op
IndexRange-8           29.8µs ± 3%
CopyRange-8            71.1µs ± 3%
IndexRangeEveryByte-8   564µs ± 2%
CopyRangeEveryByte-8    365µs ± 2%
```
and with `1048576 bytes` (1MB) of data per `largeStruct`
```
name                   time/op
IndexRange-8           27.8ms ± 3%
CopyRange-8            56.9ms ± 3%
IndexRangeEveryByte-8   154ms ± 1%
CopyRangeEveryByte-8    135ms ± 6%
```
It seems that copy performs worse when just ranging over the the struct,
but performs better when diving and iterating over the data itself.
