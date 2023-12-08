package main

import (
	"testing"
)

func TestCountCards(t *testing.T) {
	cards := "11KQQ"
	cardCount := CountCards(cards)
	expected := 3

	if len(cardCount) != expected {
		t.Errorf("expected %d, got %d", expected, len(cardCount))
	}
}

func TestGetHighCard(t *testing.T) {
	tests := []struct {
		name string
		hand string
	}{
		{
			name: "11KQQ: K",
			hand: "11KQQ",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			card, cardValue := GetHighCard(test.hand)

			if card != "K" && cardValue != 13 {
				t.Errorf("expecting card K with value 13, got %s with value %d", card, cardValue)
			}
		})
	}
}

func TestHandType(t *testing.T) {
	tests := []struct {
		name string
		hand string
		rank int
	}{
		{
			name: "Five of a kind",
			hand: "AAAAA",
			rank: 6,
		},
		{
			name: "Four of a kind",
			hand: "AKKKK",
			rank: 5,
		},
		{
			name: "Full house",
			hand: "22333",
			rank: 4,
		},
		{
			name: "Three of a kind",
			hand: "QJ777",
			rank: 3,
		},
		{
			name: "Two pair",
			hand: "11223",
			rank: 2,
		},
		{
			name: "One pair",
			hand: "12234",
			rank: 1,
		},
		{
			name: "High card",
			hand: "12345",
			rank: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cardCount := CountCards(test.hand)
			hand, rank := HandType(cardCount)

			if hand != test.name && rank != test.rank {
				t.Errorf("expected %v with value %d, got %v with value %d", test.name, test.rank, hand, rank)
			}
		})
	}
}

func TestSettleTie(t *testing.T) {
	t.Run("Check which hand is stronger", func(t *testing.T) {
		handOne := &Hand{
			Cards: "13KKQQ",
			Rank:  0,
		}

		handTwo := &Hand{
			Cards: "12KKQQ",
			Rank:  0,
		}

		handOne.GetRank()
		handTwo.GetRank()
		handOne.SettleTie(handTwo)
	})
}
