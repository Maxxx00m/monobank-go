package opened

import (
	"encoding/json"
	"github.com/asstroneer/monobank-go/pkg/base"
	"net/http"
)

type PublicClient struct {
	base.Client
}

func NewPublicClient() *PublicClient {
	return &PublicClient{
		Client: base.Client{
			HttpClient: &http.Client{},
		},
	}
}

func (c *PublicClient) GetCurrencies() (*[]Currency, error) {
	responseBody, err := c.Client.Get("bank/currency", nil)
	if err != nil {
		return nil, err
	}

	var currency []Currency
	err = json.Unmarshal(responseBody, &currency)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}
