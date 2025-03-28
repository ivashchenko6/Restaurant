package restaurant

import (
	"fmt"

	"example.com/restaurant-proj/order"
)

type Restaurant struct {
	RestaurantName string

	Orders       []order.Order
	Menu         map[string]float64
	FoodQuantity map[string]int
}

func (restaurant *Restaurant) AddOrder() error {

	//var order order.Order //New Order

	return nil
}

func (restaurant Restaurant) ShowMenu() {
	fmt.Println("Menu: ")
	fmt.Println("----------------")
	for item, price := range restaurant.Menu {
		isAvailable := restaurant.isDishAvailable(item)
		fmt.Printf("%s [%s] - %.2f$\n", item, isAvailable, price)
	}
	fmt.Println("----------------")
}

func (restaurant Restaurant) isDishAvailable(dishName string) string {
	quantity, ok := restaurant.FoodQuantity[dishName]

	if !ok || quantity < 1 {
		return "Out of Stock"
	}

	if quantity > 0 {
		return "In Stock"
	}
	return "In Stock"
}
