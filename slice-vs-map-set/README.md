## Slice vs map sets and maps

The primitive `map` type in Golang utilizes `runtime.hash64` to compute
the hash bucket from the key value. This incurs a non-trivial overhead. If a
mathematical set or map is required, it may be more performant to use slice
instead and do insertions and lookups by iteration for a sufficiently small
domain.

## Benchmarks

### Slice set

```go
var set []foo
var elem foo
...
// Insertion
isNew := true
for i, e := range set {
  if e.key == elem.key {
    isNew = false
    break
  }
}
if isNew {
  set = append(set, elem)
}
...
// Lookup
for i, e := range set {
  if e.key == desiredKey {
    return e
  }
}
```
Deletion from a slice set can lead to fragmentation and excessive GC cycles.

### Map set

```go
lookup := make(map[...]foo)
var elem foo
...
// Insertion
lookup[elem.key] = elem
...
// Lookup
return lookup[desiredKey]
```

## Flags

```
  --set-size        The number of elements inserted into the set.
  --max-elem        The maximum element in the set (min is 1). This dictates
                    approximately how often the slice set will have to
                    iterate until the end for elements not in the set.
```

## Results

### Go 1.9

With `--set-size=50`, `--max-elem=100` and 10 trials
```
name        time/op
SliceSet-8  5.78µs ± 6%
HashSet-8   6.95µs ± 3%
```
Slice sets and maps do indeed outperform `map` sets for sufficiently small sets.
