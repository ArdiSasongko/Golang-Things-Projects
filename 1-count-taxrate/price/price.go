package price

import (
	"calculator-cli/convert"
	"calculator-cli/filemanager"
	"fmt"
)

type TaxIncludePriceJob struct {
	FileManager     filemanager.FileManager `json:"-"`
	TaxRate         float64                 `json:"tax_rate"`
	InputPrice      []float64               `json:"input_price"`
	TaxIncludePrice map[string]string       `json:"tax_include_price"`
}

func NewTaxIncludePriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		FileManager:     fm,
		TaxRate:         taxRate,
		InputPrice:      []float64{10000, 20000, 30000},
		TaxIncludePrice: make(map[string]string),
	}
}

func (t *TaxIncludePriceJob) Proses() {
	t.LoadData()
	result := t.TaxIncludePrice

	for _, v := range t.InputPrice {
		aterTax := v * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", v)] = fmt.Sprintf("%.2f", aterTax)
	}

	if err := t.FileManager.WriteJson(t); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (t *TaxIncludePriceJob) LoadData() {
	lines, err := t.FileManager.ReadFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prices, err := convert.StringToFloat(lines)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.InputPrice = prices
}
