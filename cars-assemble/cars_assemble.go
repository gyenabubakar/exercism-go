package cars

import "math"

const carCost = 10000
const costForGroupsOfTen = 95000

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate / 100) * float64(productionRate)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHr := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHr) / 60
}

func CalculateCost(carsCount int) uint {
	tens := math.Floor(float64(carsCount) / 10)
	ones := carsCount - (10 * int(tens))
	cost := (int(tens) * costForGroupsOfTen) + (ones * carCost)
	return uint(cost)
}
