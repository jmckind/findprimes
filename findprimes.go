package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

// FindPrimes will return primes using the simple division method.
func FindPrimes(min int, max int) []int {
	primes := make([]int, 0)
	for i := min; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// FindPrimesSqrt will return primes using the square root method.
func FindPrimesSqrt(min int, max int) []int {
	primes := make([]int, 0)
	for i := min; i <= max; i++ {
		if isPrimeSqrt(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// FindPrimesSOE will print primes using the Sieve of Eratosthenes method.
func FindPrimesSOE(max int) []int {
	primes := make([]int, 0)
	f := make([]bool, max)
	for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
		if f[i] == false {
			for j := i * i; j < max; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < max; i++ {
		if f[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func isPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	var min = flag.Int("min", 1, "minimum value")
	var max = flag.Int("max", 100, "maximum value")
	var method = flag.String("method", "soe", "prime calculation method; \"simple\", \"sqrt\" or \"soe\"")
	flag.Parse()

	var primes []int
	start := time.Now()

	switch *method {
	case "simple":
		primes = FindPrimes(*min, *max)
	case "sqrt":
		primes = FindPrimesSqrt(*min, *max)
	case "soe":
		primes = FindPrimesSOE(*max)
	default:
		fmt.Printf("Invalid method: %s", *method)
	}

	end := time.Now()
	fmt.Printf("%s: %v: %d primes\n", *method, end.Sub(start), len(primes))
}
