package main

import (
	"fmt"
	"strings"

	"frontendmasters.com/go/cryptomasters/api"
)

func main() {
	rate, err := api.GetRate("btc")

	if err == nil {
		fmt.Printf("Currency: %v, Price: $%.2f\n", strings.ToUpper(rate.Currency), rate.Price)
	}

}
