package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(slice []string, avgPrepTime int) (estimatePrepTime int) {
	estimatePrepTime = len(slice) * avgPrepTime
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	numNoodles,numSauce := 0,0
	for _, val := range layers {
		if val == "noodles" {
			numNoodles++
		}
		if val =="sauce"{

		}
	}

	noodles = numNoodles * 50
	sauce=
	return
}

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
