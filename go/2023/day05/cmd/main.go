package main

import (
	"fmt"
	"helpers"
	"slices"
	"strings"
)

type Almanac struct {
	MapCollections []MapCollection
	Seeds          []int
}

type MapCollection struct {
	Name string
	Maps []Map
}

type Map struct {
	Destination int
	Source      int
	Range       int
}

func WriteMap(input string) Map {
	dstFromInput := strings.Split(input, " ")[0]
	srcFromInput := strings.Split(input, " ")[1]
	rngFromInput := strings.Split(input, " ")[2]

	m := Map{
		Destination: helpers.Stoi(dstFromInput),
		Source:      helpers.Stoi(srcFromInput),
		Range:       helpers.Stoi(rngFromInput),
	}

	return m
}

func NewMap(input []string) MapCollection {
	mc := MapCollection{
		Name: strings.Split(input[0], " ")[0],
	}

	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			continue
		}

		mc.Maps = append(mc.Maps, WriteMap(input[i]))
	}

	return mc
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

func CollectSeeds(input []string) []int {
	goodSeeds := []int{}
	seeds := string(input[0])

	rawSeeds := strings.Split(seeds, ":")[1]
	processedSeeds := strings.Split(rawSeeds, " ")

	for _, seed := range processedSeeds {
		if seed != "" {
			goodSeeds = append(goodSeeds, helpers.Stoi(seed))
		}
	}

	return goodSeeds
}

func NewAlmanac(input []string) *Almanac {
	a := &Almanac{
		Seeds: CollectSeeds(input),
	}

	collection := []string{}

	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			collection = RemoveEmpty(collection)
			a.MapCollections = append(a.MapCollections, NewMap(collection))
			collection = []string{}
		}

		collection = append(collection, input[i])
	}

	return a
}

func (m *Map) LocateSeed(seed int) (int, bool) {
	if seed < m.Source {
		return seed, false
	}

	if seed > m.Source+m.Range-1 {
		return seed, false
	}

	location := seed - m.Source
	foundSeed := m.Destination + location

	return foundSeed, true
}

func (a *Almanac) GetSeedLocation(seed int) int {
	found := false
	status := map[string]bool{
		"seed-to-soil":            false,
		"soil-to-fertilizer":      false,
		"fertilizer-to-water":     false,
		"water-to-light":          false,
		"light-to-temperature":    false,
		"temperature-to-humidity": false,
		"humidity-to-location":    false,
	}

	for _, mc := range a.MapCollections {
		for _, mm := range mc.Maps {
			if status[mc.Name] {
				break
			}

			seed, found = mm.LocateSeed(seed)

			if found {
				status[mc.Name] = found
			}
		}
	}

	return seed
}

func main() {
	input := helpers.ReadFile("input.txt")

	// Part 1.
	locations := []int{}

	almanac := NewAlmanac(input)
	for _, seed := range almanac.Seeds {
		locations = append(locations, almanac.GetSeedLocation(seed))
	}
	fmt.Println("Part 1: ", slices.Min(locations))
}
