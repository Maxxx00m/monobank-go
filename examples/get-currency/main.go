package main

import (
	"fmt"
	"github.com/asstroneer/monobank-go/pkg/opened"
)

func main() {
	client := opened.NewPublicClient()
	currencies, err := client.GetCurrencies()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	for _, currency := range *currencies {
		currencyA, err := currency.GetCurrencyA()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		currencyB, err := currency.GetCurrencyB()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		date, err := currency.GetDate()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("%s to %s buy %f, sell: %f, %s\n", currencyA, currencyB, currency.RateBuy, currency.RateSell, date.Format("2006-01-02 15:04:05"))
	}
}
