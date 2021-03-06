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
	return rand.Intn(1, 21)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
