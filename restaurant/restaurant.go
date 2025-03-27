package restaurant

import "example.com/restaurant-proj/order"

type Restaurant struct {
	RestaurantName string

	Orders       []order.Order
	Menu         map[string]float64
	FoodQuantity map[string]int
}
