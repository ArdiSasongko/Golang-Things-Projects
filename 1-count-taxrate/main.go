package main

import (
	"calculator-cli/filemanager"
	"calculator-cli/price"
	"fmt"
)

func main() {
	taxRates := []float64{11, 12}

	for _, t := range taxRates {
		fm := filemanager.New("./data.txt", fmt.Sprintf("./result/tax_%.0f%%.json", t))
		priceJob := price.NewTaxIncludePriceJob(fm, t)
		priceJob.Proses()
	}
}
