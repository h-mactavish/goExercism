package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(slice []string, avgPrepTime int) (estimatePrepTime int) {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	estimatePrepTime = len(slice) * avgPrepTime
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	numNoodles, numSauce := 0, 0
	for _, val := range layers {
		if val == "noodles" {
			numNoodles++
		}
		if val == "sauce" {
			numSauce++
		}
	}

	noodles = numNoodles * 50
	sauce = float64(numSauce) * 0.2
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(twoPortion []float64, numOfPortions int) (desiredPortions []float64) {

	for _, val := range twoPortion {
		desiredPortions = append(desiredPortions, val/2*float64(numOfPortions))
	}
	return
}
