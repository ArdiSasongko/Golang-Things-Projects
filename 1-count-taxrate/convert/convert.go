package convert

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var resultFloat []float64

	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float64")
		}

		resultFloat = append(resultFloat, floatPrice)
	}

	return resultFloat, nil
}
