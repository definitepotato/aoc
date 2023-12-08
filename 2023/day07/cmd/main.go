package main

import (
	"fmt"
	"helpers"
	"sort"
	"strings"
)

func CountCards(cards string) map[string]int {
	cardCount := map[string]int{}
	for _, card := range cards {
		cardCount[string(card)] += 1
	}

	return cardCount
}

func GetHighCard(cards string) (string, int) {
	cardValue := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	highCard := ""
	highCardValue := 0
	for _, card := range cards {
		if cardValue[string(card)] > highCardValue {
			highCard = string(card)
			highCardValue = cardValue[string(card)]
		}
	}

	return highCard, highCardValue
}

type Hand struct {
	Cards string
	Rank  int
	Bid   int
}

func (h *Hand) SettleTie(hand *Hand) {
	cardValue := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	for i := 0; i < len(h.Cards); i++ {
		if cardValue[string(h.Cards[i])] == cardValue[string(hand.Cards[i])] {
			continue
		}

		if cardValue[string(h.Cards[i])] > cardValue[string(hand.Cards[i])] {
			h.Rank += 1
			return
		}

		hand.Rank += 1
	}
}

func (h *Hand) GetRank() int {
	cardCount := CountCards(h.Cards)
	countBox := []int{}
	for _, count := range cardCount {
		countBox = append(countBox, count)
	}
	sort.Ints(countBox)

	for i := len(countBox) - 1; i >= 0; i-- {
		if countBox[i] == 5 {
			h.Rank = 6
			return h.Rank
		}

		if countBox[i] == 4 {
			h.Rank = 5
			return h.Rank
		}

		if countBox[i] == 3 {
			if i-1 >= 0 {
				if countBox[i-1] == 2 {
					h.Rank = 4
					return h.Rank
				}

				h.Rank = 3
				return h.Rank
			}
		}

		if countBox[i] == 2 {
			if i-1 >= 0 {
				if countBox[i-1] == 2 {
					h.Rank = 2
					return h.Rank
				}
			}

			h.Rank = 1
			return h.Rank
		}
	}

	h.Rank = 0
	return h.Rank
}

func HandType(cardCount map[string]int) (string, int) {
	countBox := []int{}
	for _, count := range cardCount {
		countBox = append(countBox, count)
	}
	sort.Ints(countBox)

	for i := len(countBox) - 1; i >= 0; i-- {
		if countBox[i] == 5 {
			return "Five of a kind", 6
		}

		if countBox[i] == 4 {
			return "Four of a kind", 5
		}

		if countBox[i] == 3 {
			if i-1 >= 0 {
				if countBox[i-1] == 2 {
					return "Full house", 4
				}

				return "Three of a kind", 3
			}
		}

		if countBox[i] == 2 {
			if i-1 >= 0 {
				if countBox[i-1] == 2 {
					return "Two pair", 2
				}
			}

			return "One pair", 1
		}
	}

	return "High card", 0
}

func NewHands(input []string) *[]Hand {
	hands := []Hand{}
	for i := 0; i < len(input); i++ {
		hand := Hand{}
		cards := strings.Split(input[i], " ")[0]
		bid := strings.Split(input[i], " ")[1]

		hand.Cards = cards
		hand.Bid = helpers.Stoi(bid)
		hand.GetRank()

		hands = append(hands, hand)
	}

	return &hands
}

func main() {
	input := helpers.ReadFile("test.txt")
	hands := NewHands(input)

	fmt.Println(hands)
}
