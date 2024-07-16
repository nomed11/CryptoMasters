package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/go/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var cryptoRate datatypes.Rate
		err = json.Unmarshal(bodyBytes, &cryptoRate)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	returnedRate := datatypes.Rate{Currency: currency, Price: 20}
	return &returnedRate, nil
}
