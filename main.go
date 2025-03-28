package main

import (
	"example.com/restaurant-proj/order"
	"example.com/restaurant-proj/restaurant"
	"example.com/restaurant-proj/utils"

	"sync"
)

func main() {
	var wg sync.WaitGroup
	menuChan := make(chan utils.Menu, 1)
	foodQuantityChan := make(chan utils.FoodQuantity, 1)
	wg.Add(2)
	go utils.LoadMenu(menuChan, &wg)
	go utils.ReceiveFoodQuantity(foodQuantityChan, &wg)

	wg.Wait()

	menu := <-menuChan
	foodQuantity := <-foodQuantityChan

	cheeseCakeFactory := restaurant.Restaurant{
		RestaurantName: "Cheese Cake Factory",
		Orders:         []order.Order{},
		Menu:           menu,
		FoodQuantity:   foodQuantity,
	}

	cheeseCakeFactory.ShowMenu()
}
