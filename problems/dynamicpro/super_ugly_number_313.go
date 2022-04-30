package dynamicpro

import (
	"math"
)

// A super ugly number is a positive integer whose prime factors are in the array primes.

// Given an integer n and an array of integers primes, return the nth super ugly number.

// The nth super ugly number is guaranteed to fit in a 32-bit signed integer.

// Example 1:

// Input: n = 12, primes = [2,7,13,19]
// Output: 32
// Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 super ugly numbers given primes = [2,7,13,19].
// Example 2:

// Input: n = 1, primes = [2,3,5]
// Output: 1
// Explanation: 1 has no prime factors, therefore all of its prime factors are in the array primes = [2,3,5].

func NthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	pointer := make([]int, len(primes))
	dp[0] = 1
	for i := 1; i <= n; i++ { // n
		min := math.MaxInt32
		// for j := 0; j < len(primes); j++ { //k
		// 	for k := 0; k < i; k++ {
		// 		if primes[j]*dp[k] > dp[i-1] { // n
		// 			min = minInt(min, primes[j]*dp[k])
		// 		}
		// 	}
		// }
		dp[i] = min
		for k := 0; k < len(primes); k++ {
			min = minInt(primes[k]*dp[pointer[k]], min)
		}
		dp[i] = min
		for k := 0; k < len(primes); k++ {
			if min == primes[k]*dp[pointer[k]] {
				pointer[k]++
			}
		}
	}
	//fmt.Println(dp)
	return dp[n-1]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
