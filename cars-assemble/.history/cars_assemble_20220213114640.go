package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
	//panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := productionRate / 60 * int(successRate/100.0)
	return int(result)
	//panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	result := carsCount/10*95000 + carsCount%10*10000
	return uint(result)
	//panic("CalculateCost not implemented")
}
