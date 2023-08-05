# monobank-go

API client for [Monobank](https://api.monobank.ua/docs/) written in Go.

## Basic usage

```go
package main

import (
	"fmt"
	"github.com/asstroneer/monobank-go/pkg/opened"
)

func main() {
	client := opened.NewPublicClient()
	currencies, err := client.GetCurrencies()
	if err != nil {
		panic(err)
	}
	fmt.Println(currencies)
}

```

More details you can find in [examples](examples) directory.