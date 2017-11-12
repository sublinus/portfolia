package main

import (
	"fmt"
	"time"
)

func main() {
	portfolia := Portfolio{}
	pos1 := Position{numOfShares: 30, symbol: "SNAP", timeOfPurchase: time.Now(), priceOfPurchase: 18.2533, name: "Snap Inc."}
	pos2 := Position{numOfShares: 10, symbol: "SBUX", timeOfPurchase: time.Now(), priceOfPurchase: 35.986272, name: "Starbucks Corp"}
	portfolia.addPosition(pos1)
	portfolia.addPosition(pos2)
	portfolia.rmPosition(0)
	fmt.Println(portfolia.positions)
	val, err := portfolia.value()
	if err != nil {
		fmt.Printf("Didn't work because: %s", err)
	}
	fmt.Printf("Current Value of Portfolio: %.4f\n", val)

}
