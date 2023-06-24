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

	// Create the map for the DrinkPrices Struct
	drinkPricesMap := make(map[string]DrinkPrices)

	for _, p := range prices {
		drinkPrices := p.Price
		drinkPrices.Name = p.Drink
		drinkPricesMap[p.Drink] = drinkPrices

	}

	// Create a map to track if we've already added the drink
	addedDrinks := make(map[string]bool)

	// Create a slice to store the filled DrinkPrices structs
	filledDrinkPrices := make([]DrinkPrices, 0)

	for _, order := range orders {
		name := order.Drink

		// Check if the drink has been added against our Map
		// Add the drink if it hasn't been

		if !addedDrinks[name] {
			drinkPrices := drinkPricesMap[name]

			filledDrinkPrices = append(filledDrinkPrices, drinkPrices)

			addedDrinks[name] = true

			/*
				fmt.Println("Drink:", name)
				//fmt.Println("Size:", size)
				fmt.Println("Small:", drinkPrices.Small)
				fmt.Println("Medium:", drinkPrices.Medium)
				fmt.Println("Large:", drinkPrices.Large)
				fmt.Println("Huge:", drinkPrices.Huge)
				fmt.Println("Mega:", drinkPrices.Mega)
				fmt.Println("Ultra:", drinkPrices.Ultra)
				fmt.Println("-------------------")
			*/
		}

	}
	// Convert filledDrinkPrices slice to JSON string
	jsonString, err := json.MarshalIndent(filledDrinkPrices, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal filledDrinkPrices to JSON:", err)
		return
	}

	// Print the JSON string to the console
	fmt.Println("Filled DrinkPrices struct:")
	fmt.Println(string(jsonString))

}
