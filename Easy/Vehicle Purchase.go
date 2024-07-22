package Easy

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	} else {
		return fmt.Sprintf("%s is clearly the better choice.", option2)
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		price := 80.0 / 100.0 // 80% do preço original
		return originalPrice * price
	} else if age < 10 {
		price := 70.0 / 100.0 // 70% do preço original
		return originalPrice * price
	} else {
		price := 50.0 / 100.0 // 50% do preço original
		return originalPrice * price
	}
}
