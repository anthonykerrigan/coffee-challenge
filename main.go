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

	type UserSummary struct {
		User         string  `json:"user"`
		OrderTotal   float32 `json:"order_total"`
		PaymentTotal float32 `json:"payment_total, omitempty"`
		Balance      float32 `json:"balance, omitempty"`
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
	userSummaries := make(map[string]float32)

	for _, order := range ordersWithPrice {
		user := order.User
		price := order.Price
		userSummaries[user] += price
	}

	summarySlice := make([]UserSummary, 0)

	for user, OrderTotal := range userSummaries {
		summary := UserSummary{
			User:         user,
			OrderTotal:   OrderTotal,
			PaymentTotal: 0,
			Balance:      0,
		}
		summarySlice = append(summarySlice, summary)
	}

	for i := range summarySlice {
		for _, payment := range payments {
			if summarySlice[i].User == payment.User {
				summarySlice[i].PaymentTotal += payment.Amount
			}
		}
		summarySlice[i].Balance = summarySlice[i].OrderTotal - summarySlice[i].PaymentTotal
	}

	// Convert the summarySlice to JSON string
	jsonString, err := json.MarshalIndent(summarySlice, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal user summaries to JSON:", err)
		return
	}

	// Print the JSON string to the console
	fmt.Println("User summaries:")
	fmt.Println(string(jsonString))
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Payments File")
	fmt.Println(payments)
}
