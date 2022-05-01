package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())

}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	//rand.Seed(time.Now().UnixNano())

	return rand.Intn(21)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
