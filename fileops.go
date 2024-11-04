package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to read vale from file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parsed stored value in file.")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) bool {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return false
	}
	return true
}
