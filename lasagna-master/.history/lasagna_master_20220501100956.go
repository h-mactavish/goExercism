package lasagna

import "strings"

// TODO: define the 'PreparationTime()' function
func PreparationTime(slice []string, avgPrepTime int) (estimatePrepTime int) {
	estimatePrepTime = len(slice) * avgPrepTime
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	noodles = strings.Count(layers, "noodles")

	return
}

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
