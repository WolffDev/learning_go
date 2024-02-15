package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	var valueText = fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	var data, err = os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 100, errors.New("failed to parse stored value value")
	}

	return value, nil
}
