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
		drinkPricesMap[p.Drink] = p.Price
	}

	for _, order := range orders {
		name := order.Drink
		size := order.Size

		drinkPrices := drinkPricesMap[name]

		fmt.Println("Drink:", name)
		fmt.Println("Size:", size)
		fmt.Println("Small:", drinkPrices.Small)
		fmt.Println("Medium:", drinkPrices.Medium)
		fmt.Println("Large:", drinkPrices.Large)
		fmt.Println("Huge:", drinkPrices.Huge)
		fmt.Println("Mega:", drinkPrices.Mega)
		fmt.Println("Ultra:", drinkPrices.Ultra)
		fmt.Println("-------------------")
	}

}
