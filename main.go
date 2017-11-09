package main

import (
	"fmt"
	"time"
)

func main() {
	pos := Position{numOfShares: 30, symbol: "SNAP", timeOfPurchase: time.Now(), priceOfPurchase: 18.2533, name: "Snap Inc."}
	rev, revPercent, err := pos.revenue()
	if err != nil {
		fmt.Printf("didn't work because: %s", err)
	}
	fmt.Printf("%.4f %.4f\n", rev, revPercent)
}
