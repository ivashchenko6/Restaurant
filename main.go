package main

import (
	"fmt"

	"example.com/restaurant-proj/order"
	"example.com/restaurant-proj/restaurant"
	"example.com/restaurant-proj/utils"

	"sync"
)

func main() {
	var wg sync.WaitGroup
	menuChan := make(chan utils.Menu)
	foodQuantityChan := make(chan utils.FoodQuantity)
	wg.Add(2)
	go utils.LoadMenu(menuChan, &wg)
	go utils.ReceiveFoodQuantity(foodQuantityChan, &wg)

	cheeseCakeFactory := restaurant.Restaurant{
		RestaurantName: "Cheese Cake Factory",
		Orders:         []order.Order{},
		Menu:           <-menuChan,
		FoodQuantity:   <-foodQuantityChan,
	}

	fmt.Println(cheeseCakeFactory)
}
