package main

import (
	"fmt"
	"helpers"
	"strings"
)

type Games struct {
	Games []Game
}

type Game struct {
	Set []Set
	Id  int
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (game *Game) IsPossible() bool {
	red, blue, green := 0, 0, 0

	for _, set := range game.Set {
		red += set.Red
		blue += set.Blue
		green += set.Green

		if set.Red > 12 {
			return false
		}

		if set.Blue > 14 {
			return false
		}

		if set.Green > 13 {
			return false
		}
	}

	return true
}

func GetGameId(game string) int {
	gameTitle := strings.Split(game, ":")[0]
	gameNumber := strings.Split(gameTitle, " ")[1]

	return helpers.Stoi(gameNumber)
}

func NewGame(game string) Game {
	return Game{
		Id:  GetGameId(game),
		Set: NewSet(game),
	}
}

func NewSet(game string) []Set {
	games := strings.Split(game, ":")[1]
	gameSets := strings.Split(games, ";")
	set := []Set{}

	for _, g := range gameSets {
		cubes := strings.Split(g, ",")
		s := Set{}

		for _, cube := range cubes {
			cubeCount := strings.Split(cube, " ")[1]
			color := strings.Split(cube, " ")[2]

			if color == "blue" {
				s.Blue = helpers.Stoi(cubeCount)
			}

			if color == "red" {
				s.Red = helpers.Stoi(cubeCount)
			}

			if color == "green" {
				s.Green = helpers.Stoi(cubeCount)
			}
		}
		set = append(set, s)
	}

	return set
}

func main() {
	input := helpers.ReadFile("input.txt")

	// Part 1.
	sum := 0
	for _, game := range input {
		g := NewGame(game)
		if g.IsPossible() {
			sum += g.Id
		}
	}
	fmt.Println(sum)
}
