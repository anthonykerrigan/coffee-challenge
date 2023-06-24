package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	pricesFile, err := ioutil.ReadFile("./data/prices.json")
	if err != nil {
		panic(err)
	}

	var prices []Prices
	err = json.Unmarshal(pricesFile, &prices)
	if err != nil {
		panic(err)
	}

	ordersFile, err := ioutil.ReadFile("./data/orders.json")
	if err != nil {
		panic(err)
	}

	var orders []Orders
	err = json.Unmarshal(ordersFile, &orders)
	if err != nil {
		panic(err)
	}

	paymentsFile, err := ioutil.ReadFile("./data/payments.json")
	if err != nil {
		panic(err)
	}

	var payments []Payments
	err = json.Unmarshal(paymentsFile, &payments)
	if err != nil {
		panic(err)
	}
	// Create the map for the DrinkPrices Struct
	drinkPricesMap := make(map[string]DrinkPrices)

	for _, p := range prices {
		drinkPrices := p.Price
		drinkPrices.Drink = p.Drink
		drinkPricesMap[p.Drink] = drinkPrices

	}

	type OrderWithPrice struct {
		Orders
		Price float32 `json:"price,omitempty"`
	}

	// Create a slice to store the combined order and price information
	ordersWithPrice := make([]OrderWithPrice, 0)

	// Fill OrderWithPrice struct for each order
	for _, order := range orders {
		drinkName := order.Drink
		size := order.Size

		// Check if the drink exists in the drinkPricesMap
		if drinkPrices, ok := drinkPricesMap[drinkName]; ok {
			// Check if the size exists in the drinkPrices struct
			var price float32
			switch size {
			case "small":
				price = drinkPrices.Small
			case "medium":
				price = drinkPrices.Medium
			case "large":
				price = drinkPrices.Large
			case "huge":
				price = drinkPrices.Huge
			case "mega":
				price = drinkPrices.Mega
			case "ultra":
				price = drinkPrices.Ultra
			}

			// Create the OrderWithPrice struct and append it to the slice
			orderWithPrice := OrderWithPrice{
				Orders: order,
				Price:  price,
			}
			ordersWithPrice = append(ordersWithPrice, orderWithPrice)
		}
	}

	// Convert filledDrinkPrices slice to JSON string
	jsonString, err := json.MarshalIndent(ordersWithPrice, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal ordersWithPrice to JSON:", err)
		return
	}

	// Print the JSON string to the console
	fmt.Println("Filled DrinkPrices struct:")
	fmt.Println(string(jsonString))

}
