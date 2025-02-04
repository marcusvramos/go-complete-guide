package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(
		investmentAmount,
		expectedReturnRate,
		years,
	)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(
	investmentAmount, 
	expectedReturnRate, 
	years float64,
) (FV float64, RFV float64) {
	FV = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	RFV = FV / math.Pow(1 + inflationRate / 100, years)
	return
}