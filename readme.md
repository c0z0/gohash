# GoHash

A simple implementation of two types of hash maps

- Open hash
- Closed hash

Currently only `string` to `int` mapping is supported

## Usage

```go
  package main

  import (
    "fmt"
    hash "github.com/c0z0/gohash/closed-hash" // or "github.com/c0z0/gohash/open-hash"
    )

  func main() {
    h := hash.CreateHash(1000) // create a hash of size 1000

    h.Put(1, "test")
    h.Put(2, "test2")
    h.Put(3, "test3")
    h.Put(4, "test4")

    fmt.Println(h.Get("test"))
    fmt.Println(h.Get("test2"))
    fmt.Println(h.Get("test3"))
    fmt.Println(h.Get("test4"))
  }
```
