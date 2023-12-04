package main

import (
	"helpers"
	"testing"
)

const CARD = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

func TestNewCard(t *testing.T) {
	got := NewCard(CARD)

	if got.Id != 1 {
		t.Errorf("got %d expected 1", got.Id)
	}
}

func TestCountPoints(t *testing.T) {
	card := NewCard(CARD)
	points := card.CountPoints()

	if points != 8 {
		t.Errorf("got %d expected 8", points)
	}
}

func TestCountCards(t *testing.T) {
	input := helpers.ReadFile("../test.txt")

	cards := []Card{}
	for _, card := range input {
		cards = append(cards, NewCard(card))
	}

	copies := CountCards(cards)

	if copies != 30 {
		t.Errorf("got %d expected 30", copies)
	}
}
