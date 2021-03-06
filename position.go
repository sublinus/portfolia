package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Position struct {
	NumOfShares     float64
	TimeOfPurchase  time.Time
	PriceOfPurchase float64
	Symbol          string
	Name            string
}

func (p Position) currentPrice() (float64, time.Time, error) {
	url := "http://dev.markitondemand.com/MODApis/Api/v2/Quote/json?symbol=" + p.Symbol
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return 0, time.Time{}, err
	}
	var stock struct {
		Price float64 `json:"LastPrice"`
		Time  string  `json:"Timestamp"`
	}
	json.NewDecoder(res.Body).Decode(&stock)
	if stock.Price != 0 && stock.Time != "" {
		layout := "Mon Jan 2 15:04:05 MST-07:00 2006"
		t, err := time.Parse(layout, stock.Time)
		if err != nil {
			return 0, time.Time{}, err
		}
		return stock.Price, t, nil
	}
	return 0, time.Time{}, fmt.Errorf("Something went wrong: Stock price: %f Timestamp: %s", stock.Price, stock.Time)
}

func (p Position) revenuePerShare() (float64, float64, error) {
	price, _, err := p.currentPrice()
	if err != nil {
		return 0, 0, err
	}
	revenue := price - p.PriceOfPurchase
	revenuePercentage := revenue / p.PriceOfPurchase

	return revenue, revenuePercentage, nil
}

func (p Position) revenue() (float64, float64, error) {
	revenue, revenuePercent, err := p.revenuePerShare()
	if err != nil {
		return 0, 0, err
	}
	return revenue * p.NumOfShares, revenuePercent * 100, nil
}
