package utils

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"example.com/restaurant-proj/table"
)

type Menu map[string]float64
type FoodQuantity map[string]int

func LoadMenu(menuChan chan Menu, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Loading the menu...")
	time.Sleep(2 * time.Second)

	fmt.Println("Menu is ready!")
	menuChan <- Menu{
		"Margherita Pizza":           12.99,
		"Pepperoni Pizza":            14.99,
		"BBQ Chicken Pizza":          15.99,
		"Caesar Salad":               9.49,
		"Greek Salad":                8.99,
		"Grilled Salmon":             18.99,
		"Steak with Mashed Potatoes": 22.49,
		"Spaghetti Carbonara":        13.99,
		"Chicken Alfredo Pasta":      14.99,
		"Garlic Bread":               4.99,
		"French Fries":               3.99,
		"Chocolate Cake":             6.99,
		"Cheesecake":                 7.49,
		"Apple Pie":                  5.99,
		"Cappuccino":                 4.49,
		"Latte":                      4.99,
		"Fresh Orange Juice":         3.99,
		"Mineral Water":              2.49,
	}

	close(menuChan)
}

func ReceiveFoodQuantity(foodQuantityChan chan FoodQuantity, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Loading FoodQuantity")

	time.Sleep(2 * time.Second)

	fmt.Println("Food Quantity is ready!")
	foodQuantityChan <- FoodQuantity{
		"Margherita Pizza":           10,
		"Pepperoni Pizza":            8,
		"BBQ Chicken Pizza":          6,
		"Caesar Salad":               12,
		"Greek Salad":                10,
		"Grilled Salmon":             5,
		"Steak with Mashed Potatoes": 4,
		"Spaghetti Carbonara":        7,
		"Chicken Alfredo Pasta":      9,
		"Garlic Bread":               15,
		"French Fries":               20,
		"Chocolate Cake":             6,
		"Cheesecake":                 8,
		"Apple Pie":                  10,
		"Cappuccino":                 12,
		"Latte":                      10,
		"Fresh Orange Juice":         15,
		"Mineral Water":              25,
	}
	close(foodQuantityChan)
}

func ReadMultipleWords(strPtr *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*strPtr = scanner.Text()
}

func GetTables(tablesChan chan []table.Table, wg *sync.WaitGroup) {
	defer wg.Done()
	tablesChan <- []table.Table{
		{TableID: 16, PeopleQuantity: 5, ReceiptAmount: 0, IsAvailable: false},
		{TableID: 12, PeopleQuantity: 0, ReceiptAmount: 0, IsAvailable: true},
		{TableID: 19, PeopleQuantity: 2, ReceiptAmount: 0, IsAvailable: false},
		{TableID: 65, PeopleQuantity: 1, ReceiptAmount: 0, IsAvailable: false},
	}

	close(tablesChan)
}
