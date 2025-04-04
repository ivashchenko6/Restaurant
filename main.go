package main

import (
	"fmt"

	"example.com/restaurant-proj/order"
	"example.com/restaurant-proj/restaurant"
	"example.com/restaurant-proj/table"
	"example.com/restaurant-proj/utils"

	"sync"
)

func main() {

	//Loading menu:
	//----------
	menu, foodQuantity, tables := loadServices()
	//----------
	restaurant := restaurant.Restaurant{
		RestaurantName: "Cheese Cake Factory",
		Orders:         []order.Order{},
		Tables:         tables,
		Menu:           menu,
		FoodQuantity:   foodQuantity,
	}
	table16 := table.Table{TableID: 16, Orders: make(map[string]int), PeopleQuantity: 5, ReceiptAmount: 0, IsAvailable: false}
	// table12 := table.Table{TableID: 12, Orders: make(map[string]int), PeopleQuantity: 0, ReceiptAmount: 0, IsAvailable: true}
	// table19 := table.Table{TableID: 19, Orders: make(map[string]int), PeopleQuantity: 2, ReceiptAmount: 0, IsAvailable: false}
	// table65 := table.Table{TableID: 65, Orders: make(map[string]int), PeopleQuantity: 1, ReceiptAmount: 0, IsAvailable: false}

	fmt.Println("Table:", table16)
	restaurant.AddOrder(&table16)

	fmt.Println("Table:", table16)

	//TODO: Make an system with Tables, All orders by one  table, total amount $$$ per visit (per table)
	//TODO: Ask how many servings customer would like per each  dish
}

func loadServices() (utils.Menu, utils.FoodQuantity, []table.Table) {
	var wg sync.WaitGroup
	menuChan := make(chan utils.Menu, 1)
	foodQuantityChan := make(chan utils.FoodQuantity, 1)
	tablesChan := make(chan []table.Table, 1)
	wg.Add(3)
	go utils.LoadMenu(menuChan, &wg)
	go utils.ReceiveFoodQuantity(foodQuantityChan, &wg)
	go utils.GetTables(tablesChan, &wg)

	wg.Wait()

	menu := <-menuChan
	foodQuantity := <-foodQuantityChan
	tables := <-tablesChan

	return menu, foodQuantity, tables
}
