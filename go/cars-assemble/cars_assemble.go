package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successfulCars := float64(productionRate) * (successRate / 100)
	return successfulCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	prodRateMinute := float64(productionRate) / 60
	ratePerMinute := prodRateMinute * (successRate / 100)

	return int(ratePerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	batchCost := (carsCount / 10) * 95000
	spareCost := 0
	if carsCount%10 != 0 {
		spareCost = (carsCount % 10) * 10000
	}
	var totalCost uint = uint(batchCost + spareCost)
	return totalCost
}
