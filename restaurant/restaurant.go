package restaurant

import (
	"fmt"
	"strconv"

	"example.com/restaurant-proj/order"
	"example.com/restaurant-proj/table"
	"example.com/restaurant-proj/utils"
)

type Restaurant struct {
	RestaurantName string
	Tables         []table.Table
	Orders         []order.Order
	Menu           map[string]float64
	FoodQuantity   map[string]int
}

func (restaurant *Restaurant) AddOrder(table *table.Table) {
	//scanner := bufio.NewScanner(os.Stdin)
	var orderNameString string
	var order order.Order = order.Order{TableID: 0, Dishes: make(map[string]int), Check: 0.0, Status: "in progress"} //New Order
	restaurant.ShowMenu()
	fmt.Print("What would you like to order(finish `F`) ")
	for {
		fmt.Print(">>>")
		utils.ReadMultipleWords(&orderNameString)
		//scanner.Scan()
		//orderNameString = scanner.Text()

		if orderNameString == "F" {
			fmt.Println("Finished")
			break
		}

		_, exist := restaurant.Menu[orderNameString] //Does the dish name exist in the  menu

		if !exist {
			fmt.Println("Invalid Name.Try Again")
			continue
		}

		servingsAmountStr := ""
		var servingsAmount int
		for {
			fmt.Printf("How many servings would you like [%s - MAX %d]:", orderNameString, restaurant.FoodQuantity[orderNameString])
			utils.ReadMultipleWords(&servingsAmountStr)
			servingsAmountTemp, err := strconv.Atoi(servingsAmountStr)

			if err != nil {
				fmt.Println("Invalid Amount of Servings. Try again")
				continue
			}

			if servingsAmountTemp < 1 || (restaurant.FoodQuantity[orderNameString]-servingsAmountTemp) < 0 {
				fmt.Println("Servings must be positive and be less or equal to MAX numbers in stock")
				continue
			}

			servingsAmount = servingsAmountTemp
			break
		}

		restaurant.FoodQuantity[orderNameString] -= servingsAmount

		fmt.Println("Added: ", servingsAmount, orderNameString)
		//
		order.Dishes[orderNameString] += servingsAmount // For restaurant
		table.Orders[orderNameString] += servingsAmount // For table
		//

	}

	fmt.Println("Total Order")
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

func (restaurant Restaurant) tableExist(tableId int) bool {
	for _, currentTable := range restaurant.Tables {
		if currentTable.TableID == tableId {
			return true
		}
	}
	return false
}
