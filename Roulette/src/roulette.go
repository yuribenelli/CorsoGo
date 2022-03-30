package main

import (
	"fmt"
	"math/rand"
)

var playAgain bool = true
var currMoney int = 100

func main() {
	fmt.Printf("\nBenvenuto nel gioco della roulette\n")
	for playAgain {
		bet := 0
		canBet := true
		var betNumber uint
		fmt.Printf("La somma che possiedi è %v . \n Quanto vuoi puntare?\n", currMoney)
		fmt.Scan(&bet)

		// Controllo input
		if bet > currMoney {
			fmt.Println("Non hai abbastanza soldi.\n", currMoney)
			canBet = false

		}
		if bet <= 0 {
			fmt.Printf("La scommessa deve avere,come importo, un numero maggiore di zero\n")
			canBet = false
		}
		if canBet {
			fmt.Printf("la tua scommessa è di %v \nSu quale numero vuoi puntare? \n", bet)
			fmt.Scan(&betNumber)
			if betNumber < 0 || betNumber >= 90 {
				fmt.Printf("hai inserito un numero non valido. Deve essere compreso fra 1 e 90\n")
				canBet = false
			}
			rndNumber := rand.Intn(90)
			fmt.Printf("è uscito il numero %v. Avevi puntato sul numero %v\n", rndNumber, betNumber)
			if rndNumber == int(betNumber) {
				fmt.Printf("HAI VINTO!")
				currMoney = currMoney + bet*2
				fmt.Printf("La tua vincita ammonta a %v. Adesso hai a disposizione %v crediti\n", bet*2, currMoney)
			}
			if rndNumber != int(betNumber) {
				currMoney = currMoney - bet
				fmt.Printf("Non hai vinto. Ritenta! Il tuo patrimonio è di %v crediti\n", currMoney)
			}
			fmt.Printf("Vuoi giocare ancora? y/n\n")
			goOnChoice := ""
			fmt.Scan(&goOnChoice)
			switch goOnChoice {
			case "y":
				playAgain = true
			case "n":
				playAgain = false
			}

		}
	}
}
