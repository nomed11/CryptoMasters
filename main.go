package main

import (
	"fmt"

	"frontendmasters.com/go/cryptomasters/api"
)

func main() {
	rate, err := api.GetRate("eth")

	fmt.Print(*rate, err)

}
