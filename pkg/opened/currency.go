package opened

import (
	"errors"
	"fmt"
	"github.com/rmg/iso4217"
	"strconv"
	"time"
)

type Currency struct {
	CurrencyCodeA int     `json:"currencyCodeA"`
	CurrencyCodeB int     `json:"currencyCodeB"`
	Date          int     `json:"date"`
	RateBuy       float64 `json:"rateBuy"`
	RateCross     float64 `json:"rateCross"`
	RateSell      float64 `json:"rateSell"`
}

func (c *Currency) GetCurrencyA() (string, error) {
	return c.getCurrency(c.CurrencyCodeA)
}

func (c *Currency) GetCurrencyB() (string, error) {
	return c.getCurrency(c.CurrencyCodeB)
}

func (c *Currency) GetDate() (*time.Time, error) {
	i, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(i, 0)
	return &tm, nil
}

func (c *Currency) getCurrency(code int) (string, error) {
	name, _ := iso4217.ByCode(code)
	if name == "" {
		return "", errors.New(fmt.Sprintf("currency code %d not found", code))
	}
	return name, nil
}
