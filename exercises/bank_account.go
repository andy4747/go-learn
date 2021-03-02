package exercises

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Account struct
type Account struct {
	Balance int
}

// getBalance returns the total balance of a user
func (a Account) getBalance() (int, error) {
	if a.Balance < 0 {
		fmt.Println("The total balance is 0")
		fmt.Println("----------------------")
		return a.Balance, errors.New("insufficient balance")
	}
	return a.Balance, nil
}

// Deposit adds an amount to user account
func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

// Withdraw removes some amount from user account
func (a *Account) Withdraw(amount int) {
	if a.Balance < amount {
		fmt.Println("Insufficient Balance")
		return
	}
	a.Balance -= amount
}

// BankApplication function
func BankApplication() {
	account := new(Account)
	for true {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter (1) to Deposit | Enter (2) to Withdraw | Enter (3) to view balance: ")
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if input == 1 {
			fmt.Print("Enter the amount to deposit: ")
			scanner.Scan()
			depositAmount, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			account.Deposit(int(depositAmount))
			balance, _ := account.getBalance()
			fmt.Println("The available balance is", balance)
		} else if input == 2 {
			fmt.Print("Enter the amount to withdraw: ")
			scanner.Scan()
			withdrawAmount, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			account.Withdraw(int(withdrawAmount))
			balance, error := account.getBalance()
			if error != nil {
				fmt.Println("Insufficient Balance")
				continue
			}
			fmt.Println("The available balance is", balance)
		} else if input == 3 {
			balance, _ := account.getBalance()
			fmt.Println("The available balance is", balance)
		} else {
			return
		}
	}
}
