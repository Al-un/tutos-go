package main

import "fmt"

func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func getPrimeNumbers(from int, to int) []int {
	var primes []int

	for i := from; i < to; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func getPrimeFactors(n int, maxCheckedPrimeNumber int) ([]int, int) {
	primeNumbers := getPrimeNumbers(2, maxCheckedPrimeNumber)
	var primeFactors []int

	number := n
	for _, prime := range primeNumbers {
		// fmt.Println("Checking prime ", prime)
		if number == 1 {
			number = 0
			break
		}

		for number%prime == 0 {
			primeFactors = append(primeFactors, prime)
			number /= prime
			// fmt.Println("Number is now ", number, " with factors ", primeFactors)
		}
	}

	return primeFactors, number
}

func main() {
	primeFactors, remainder := getPrimeFactors(600851475143, 10000)
	fmt.Println("Prime factors are ", primeFactors, " with remainder ", remainder)
}
