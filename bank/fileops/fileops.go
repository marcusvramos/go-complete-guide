package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func WriteFloatToFile(fileName string, value float64) {
	valueString := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueString), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("error reading file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("error parsing data")
	}

	return value, nil
}
