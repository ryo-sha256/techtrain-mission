package Utils

import "math/rand"

// generates a random number from 0.01 to 1.00
// @param n: number of random number that you want
// @returns the array of random numbers between 0.01 and 1.00
func RandGenerator(n int) []float32 {
	res := make([]float32, n)
	for i := range res {
		res[i] = 0.01 + rand.Float32()*(0.99)
	}
	return res
}
