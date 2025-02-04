package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResultToFile(eb, profit, ratio float64) {
	results := fmt.Sprintf("Earnings before tax: %.2f\nProfit: %.2f\nRatio: %.2f\n", eb, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {
	revenue, err1 := getUserInput("Enter revenue: ")
	expenses, err2 := getUserInput("Enter expenses: ")
	taxRate, err3 := getUserInput("Enter tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1, err2, err3)
		panic("Invalid input")
	}
	
	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	outputFinancial(ebt, profit, ratio)
	storeResultToFile(ebt, profit, ratio)
}

func getUserInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("must be greater than 0")
	}

	return input, nil
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func outputFinancial(ebt, profit, ratio float64) {
	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}