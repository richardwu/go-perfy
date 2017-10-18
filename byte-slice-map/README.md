# Maps with byte slice keys

Go maps do not permit using slices (with arbitrary size) as keys in maps.

For byte slices, one has a few options for re-hashing the slice into a
mappable type.

## Benchmarks

### String()

One can simply cast the byte slice into an immutate `string` type
```go
lookup := make(map[string]foo)
var key []byte
var input foo
...
lookup[string(key)] = input
...
```

### MD5

[MD5](https://en.wikipedia.org/wiki/MD5) converts any byte slice into a 128-bit
hash value (16 bytes). It is not cryptographically safe anymore, but can be used
as a robust hashing function that's fairly fast.

```go
import "crypto/md5"

lookup := make(map[[16]byte]foo)
var key []byte
var input foo
...
lookup[md5.Sum(key)] = input
...
```

### CRC32

[CRC32](https://en.wikipedia.org/wiki/Cyclic_redundancy_check) is a common
hashing scheme used for error checking. It produces a 32-bit hash value (`uint32`).
Collisions are possible, so it is advised that one check if the key is already
present in the map before updating.

```go
import "hash/crc32"

lookup := make(map[uint32]foo)
var key []byte
var input foo
...
table := crc32.MakeTable(crc32.IEEE)
lookup[crc32.Checksum(key, table)] = input
...
```

## Flags

```
  --key-size      The length of the byte slices used for the keys.
  --num-keys      The number of keys to map.
```

## Results

### Go 1.9

With `--key-size=32`, `--num-keys=30` and 10 trials
```
name         time/op
StringMap-4  77.8µs ± 2%
MD5Map-4     86.5µs ±12%
CRC32Map-4   79.4µs ± 2%
```
Directly string casting the byte slice keys seems to be the fastest option
and the hashing functions yield no benefit. Furthermore, hash functions
always have the potential of collisions (albeit __rare__ for crypto hashes).
