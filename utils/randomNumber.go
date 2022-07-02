package Utils

import (
	"math/rand"
	models "techtrain-mission/models"
)

// generates a random number between 0.01 and 1.00
// @param n: number of random number that you need
// @returns: an array of random numbers between 0.01 and 1.00
func RandGenerator(n int) []float32 {
	res := make([]float32, n)
	for i := range res {
		res[i] = 0.01 + rand.Float32()*(0.99)
	}
	return res
}

// returns one of the following rarities: common, uncommon, rare, ultraRare
// Depending on a float number that is passed into this function,
// it returns a correct rarity accordingly
func RarityTable(rarity float32, r models.RarityList) string {
	// switch statement goes here
	switch {
	case rarity >= r.Common:
		return "common"
	case r.Common > rarity && rarity >= r.Uncommon:
		return "uncommon"
	case r.Uncommon > rarity && rarity >= r.Rare:
		return "rare"
	case r.Rare > rarity:
		return "ultraRare"

	default:
		return "common"
	}
}
