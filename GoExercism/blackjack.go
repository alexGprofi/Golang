package main

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ten", "jack", "queen", "king":
		value = 10
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var action string
	dealer := ParseCard(dealerCard)
	carta1 := ParseCard(card1)
	carta2 := ParseCard(card2)
	switch {
	case carta1 == 11 && carta2 == 11:
		action = "P"
	case (carta1 + carta2) == 21:
		if dealer < 10 {
			action = "W"
		} else {
			action = "S"
		}

	case (carta1+carta2) < 21 && (carta1+carta2) >= 17:
		action = "S"
	case (carta1+carta2) <= 16 && (carta1+carta2) >= 12:
		if dealer >= 7 {
			action = "H"
		} else {
			action = "S"
		}

	case (carta1 + carta2) <= 11:
		action = "H"
	}
	return action
}

func main() {
	fmt.Println(ParseCard("ace"))
	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace"))
	fmt.Println(FirstTurn("five", "queen", "ace"))
}
