package main

import (
	"fmt"
	"helpers"
	"strings"
)

type Card struct {
	Winners []int
	Numbers []int
	Id      int
	Copies  int
}

func (card *Card) CountPoints() int {
	sum := 0

	for _, winner := range card.Winners {
		for _, number := range card.Numbers {
			if winner == number {
				if sum == 0 {
					sum += 1
				} else {
					sum *= 2
				}
			}
		}
	}

	return sum
}

func CountCards(cards []Card) int {
	winningNumbers := 0

	for i := 0; i < len(cards); i++ {
		// Calculate number of winners on card.
		for _, winner := range cards[i].Winners {
			for _, number := range cards[i].Numbers {
				if winner == number {
					winningNumbers += 1
					cards[i+winningNumbers].Copies += 1
				}
			}
		}

		// Make copies.
		for n := 0; n < cards[i].Copies; n++ {
			for j := 0; j < winningNumbers; j++ {
				cards[i+j+1].Copies += 1
			}
		}
		winningNumbers = 0
	}

	// Count copies.
	copies := 0
	for _, card := range cards {
		copies += card.Copies + 1
	}

	return copies
}

func NewCard(card string) Card {
	cardTitle := strings.Split(card, ":")[0]
	cardId := strings.Split(cardTitle, " ")[1]

	allNumbers := strings.Split(card, ":")[1]
	winningNumbers := strings.Split(allNumbers, "|")[0]
	cardNumbers := strings.Split(allNumbers, "|")[1]

	winningNumbersSlice := strings.Split(winningNumbers, " ")
	cardNumbersSlice := strings.Split(cardNumbers, " ")

	myCard := Card{}
	myCard.Id = helpers.Stoi(cardId)

	for _, num := range winningNumbersSlice {
		if num == "" {
			continue
		}
		myCard.Winners = append(myCard.Winners, helpers.Stoi(num))
	}

	for _, num := range cardNumbersSlice {
		if num == "" {
			continue
		}
		myCard.Numbers = append(myCard.Numbers, helpers.Stoi(num))
	}

	return myCard
}

func main() {
	input := helpers.ReadFile("input.txt")

	cards := []Card{}
	for _, card := range input {
		cards = append(cards, NewCard(card))
	}

	// Part 1.
	sum := 0
	for _, card := range cards {
		sum += card.CountPoints()
	}
	fmt.Println("Part 1: ", sum)

	// Part 2.
	copies := CountCards(cards)
	fmt.Println("Part 2: ", copies)
}
