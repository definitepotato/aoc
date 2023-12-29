package main

import (
	"fmt"
	"helpers"
	"strings"
)

type Races struct {
	Race []Race
}

type Race struct {
	Time     int
	Distance int
}

func RemoveEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func (r *Race) WaysToWin() int {
	ways := 0

	for timeHeld := 0; timeHeld < r.Time; timeHeld++ {
		remainingRaceTime := r.Time - timeHeld
		traveled := timeHeld * remainingRaceTime

		if traveled > r.Distance {
			ways += 1
		}
	}

	return ways
}

func NewRaces(input []string) *Races {
	allTimes := strings.Split(input[0], ":")[1]
	allDistances := strings.Split(input[1], ":")[1]

	times := strings.Split(allTimes, " ")
	distances := strings.Split(allDistances, " ")

	t := RemoveEmpty(times)
	d := RemoveEmpty(distances)

	races := &Races{}

	for i := 0; i < len(t); i++ {
		r := Race{
			Time:     helpers.Stoi(t[i]),
			Distance: helpers.Stoi(d[i]),
		}
		races.Race = append(races.Race, r)
	}

	return races
}

func NewRace(input []string) *Races {
	allTimes := strings.Split(input[0], ":")[1]
	allDistances := strings.Split(input[1], ":")[1]

	times := strings.Split(allTimes, " ")
	distances := strings.Split(allDistances, " ")

	race := Race{
		Time:     helpers.Stoi(strings.Join(times, "")),
		Distance: helpers.Stoi(strings.Join(distances, "")),
	}

	races := &Races{}
	races.Race = append(races.Race, race)

	return races
}

func main() {
	input := helpers.ReadFile("input.txt")

	// Part 1.
	races := NewRaces(input)
	winners := []int{}
	answer := 1
	for _, race := range races.Race {
		winners = append(winners, race.WaysToWin())
	}

	for _, winner := range winners {
		answer *= winner
	}
	fmt.Println(answer)

	// Part 2.
	race := NewRace(input)
	waysToWin := 0
	for _, race := range race.Race {
		waysToWin = race.WaysToWin()
	}
	fmt.Println(waysToWin)
}
