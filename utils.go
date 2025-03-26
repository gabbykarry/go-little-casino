package main

import (
	"fmt"
	"os"
)

func GetName() {
	name := ""

	fmt.Println("-----------")
	fmt.Print("Pls enter your name:")

	for {
		_, err := fmt.Scanln(&name)
		if err != nil {
			os.Exit(1)
		}
		if name != "" {
			break
		} else {
			fmt.Println("Please enter a name")
		}
	}
	fmt.Printf("Hello %s\n", name)
}

func GetBet(balance uint) uint {
	fmt.Println("---------")
	var bet uint

	for {
		fmt.Printf("Please enter your bet, 0 to quit (balance = $%d): ", balance)
		_, err := fmt.Scanln(&bet)
		if err != nil {
			_ = fmt.Errorf("exiting system %s", err)
			os.Exit(1)
		}
		if bet > balance {
			fmt.Printf("You can place a bet higher than %d\n", balance)
		} else {
			break
		}

	}
	return bet
}
