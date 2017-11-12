package main

import (
	"fmt"
)

func main() {
	portfolia := Portfolio{}
	// pos1 := Position{NumOfShares: 30, Symbol: "SNAP", TimeOfPurchase: time.Now(), PriceOfPurchase: 18.2533, Name: "Snap Inc."}
	// pos2 := Position{NumOfShares: 10, Symbol: "SBUX", TimeOfPurchase: time.Now(), PriceOfPurchase: 35.986272, Name: "Starbucks Corp"}
	// portfolia.addPosition(pos1)
	// portfolia.addPosition(pos2)
	// portfolia.rmPosition(0)
	// portfolia.addPosition(pos1)
	portfolia.importJSON("portfolio.json")
	val, err := portfolia.value()
	if err != nil {
		fmt.Printf("Didn't work because: %s", err)
	}
	fmt.Printf("Current Value of Portfolio: %.4f\n", val)
}
