package main

import (
	"fmt"
	"example.com/bank/fileops"
)

const accounBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accounBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----")
	}
	fmt.Println("Welcome to Go Bank")

	for {
		presentOptions()

		choice := int(0)
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
			fileops.WriteFloatToFile(accountBalance, accounBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accounBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}

}
