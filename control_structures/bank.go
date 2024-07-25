package main

import (
	"fmt"
	"os"
	"strconv"
)

const accounBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accounBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accounBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check the balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is ", accountBalance)
		case 2:
			fmt.Println("How much? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("How much?")
			var WithdrawAmount float64
			fmt.Scan(&WithdrawAmount)

			if WithdrawAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				// return
				continue
			}

			if WithdrawAmount > accountBalance {
				fmt.Println("Invalid Amount! You cant withdraw more than you have.")
				// return
				continue
			}
			accountBalance = accountBalance - WithdrawAmount
			fmt.Println("Balance updated! New Balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}

}
