package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	var choice int

	if err != nil {
		fmt.Println("Error ", err)
		fmt.Println("-----------------")
	}

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("Hello, ", randomdata.FirstName(0))

	for choice != 4 {
		showMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is %.2f\n", accountBalance)
		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount! Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			fmt.Print("Your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
	
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds! Cannot withdraw more than account balance.")
				continue
			}

			accountBalance -= withdrawalAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using the Go Bank!")
			return
		}
	}
}