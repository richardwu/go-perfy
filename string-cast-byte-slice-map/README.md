# Accessing string map with byte slices

As discussed in [Maps with byte slice keys](https://github.com/richardwu/go-perfy/tree/master/byte-slice-map), `string()` cast of `[]byte` keys is the most
efficient way of mapping byte slices to values.

There is [one compiler optimization](https://github.com/golang/go/blob/e97209515ad8c4042f5a3ef32068200366892fc2/src/runtime/string.go#L130-L132)
where directly indexing with a `string()`
cast of a byte slice is more efficient than casting the byte slice into a `string`
first.

## Benchmarks

### Direct indexing

```go
lookup := make(map[string]foo)
var key []byte
var input foo
...
lookup[string(key)] = input
...
```

### Cast first

```go
lookup := make(map[string]foo)
var key []byte
var input foo
...
strKey := string(key)
...
lookup[strKey] = input
...
```

## Flags

```
  --key-size      The length of the byte slices used for the keys.
  --num-keys      The number of keys to map.
```

## Results

With `--key-size=32`, `--num-keys=30` and 10 trials
```
name           time/op
DirectIndex-8  4.94µs ± 2%
CastFirst-8    5.89µs ± 2%
```
It is true that direct indexing with a `string()` cast on a `[]byte` key
performs better than casting the key into a `string` type first.
