package utils

import (
	"math"
)

// Value type
type Value int

// Key type
type Key string

func isPrime(n uint) bool {
	if n <= 1 {
		return false
	}

	var nd uint

	for d := uint(2); d < uint(math.Sqrt(float64(n))); d++ {
		if n%d == 0 {
			nd++
		}
	}

	return nd == 0
}

// FindNextPrime finds the smallest prime number bigger or equal to n
func FindNextPrime(n uint) uint {
	if isPrime(n) {
		return n
	}

	return FindNextPrime(n + 1)
}

// Hash hashes the key
func Hash(key Key, size uint) uint {
	var hashVal uint

	for _, c := range key {
		hashVal = hashVal<<5 + uint(c)
	}

	return hashVal % size
}
