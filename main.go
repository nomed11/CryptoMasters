package main

import (
	"fmt"
	"strings"
	"sync"

	"frontendmasters.com/go/cryptomasters/api"
)

func main() {
	currencies := []string{"btc", "eth", "xrp"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currency string) {
			defer wg.Done()
			getCurrencyData(currency)
		}(currency)
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("Currency: %v, Price: $%.2f\n", strings.ToUpper(rate.Currency), rate.Price)
	}
}
