package main

import (
	"fmt"

	hash "github.com/c0z0/gohash/open-hash"
)

func main() {
	h := hash.CreateHash(3)

	h.Put(1, "test")
	h.Put(2, "test2")
	h.Put(3, "test3")
	h.Put(4, "test4")

	fmt.Println(h.Get("test"))
	fmt.Println(h.Get("test2"))
	fmt.Println(h.Get("test3"))
	fmt.Println(h.Get("test4"))

}
