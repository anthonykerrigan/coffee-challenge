package main

type Prices struct {
	Drink string      `json:"drink_name"`
	Price DrinkPrices `json:"prices"`
}

type DrinkPrices struct {
	Name   string  `json:"user"`
	Small  float32 `json:"small,omitempty"`
	Medium float32 `json:"medium,omitempty"`
	Large  float32 `json:"large,omitempty"`
	Huge   float32 `json:"huge,omitempty"`
	Mega   float32 `json:"mega,omitempty"`
	Ultra  float32 `json:"ultra,omitempty"`
}

type Orders struct {
	User  string `json:"user"`
	Drink string `json:"drink"`
	Size  string `json:"size"`
}

type Payments struct {
	User   string `json:"user"`
	Amount string `json:"amount"`
}
