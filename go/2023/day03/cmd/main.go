package main

import (
	"fmt"
	"helpers"
	"strconv"
)

type Schematic struct {
	Row   []string
	Parts []Part
	Sum   int
}

type Part struct {
	Coordinates []Coordinate
	Number      int
	IsPart      bool
}

type Coordinate struct {
	Y int
	X int
}

func IsNumber(n string) bool {
	_, err := strconv.Atoi(n)

	return err == nil
}

func NewSchematic(input []string) Schematic {
	schematic := Schematic{}
	schematic.Row = append(schematic.Row, input...)

	schematic.ScanParts()
	schematic.MarkParts()

	return schematic
}

func (s *Schematic) MarkParts() {
	for y := 0; y < len(s.Row); y++ { // iterate y pos
		for x := 0; x < len(s.Row[y]); x++ { // iterate x pos
			currentPosValue := string(s.Row[y][x])

			if currentPosValue == "." || IsNumber(currentPosValue) {
				continue
			}

			for i := 0; i < len(s.Parts); i++ {
				for _, partCoordinates := range s.Parts[i].Coordinates {
					if partCoordinates.Y == y-1 && partCoordinates.X == x-1 { // up left
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y-1 && partCoordinates.X == x { // up
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y-1 && partCoordinates.X == x+1 { // up right
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y && partCoordinates.X == x-1 { // left
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y && partCoordinates.X == x+1 { // right
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y+1 && partCoordinates.X == x-1 { // down left
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y+1 && partCoordinates.X == x { // down
						s.Parts[i].IsPart = true
						break
					}

					if partCoordinates.Y == y+1 && partCoordinates.X == x+1 { // down right
						s.Parts[i].IsPart = true
						break
					}
				}
			}
		}
	}
}

func (s *Schematic) ScanParts() {
	part := Part{}
	num := ""

	for y := 0; y < len(s.Row); y++ { // iterate y pos
		for x := 0; x < len(s.Row[y]); x++ { // iterate x pos
			currentPosValue := string(s.Row[y][x])

			if IsNumber(currentPosValue) {
				coordinates := Coordinate{
					X: x,
					Y: y,
				}
				part.Coordinates = append(part.Coordinates, coordinates)
				num += currentPosValue
			}

			if !IsNumber(currentPosValue) && len(part.Coordinates) >= 1 {
				part.Number = helpers.Stoi(num)
				s.Parts = append(s.Parts, part)

				part = Part{}
				num = ""
			}
		}
	}
}

func main() {
	input := helpers.ReadFile("input.txt")
	schematic := NewSchematic(input)

	sum := 0
	for _, part := range schematic.Parts {
		fmt.Println(part)
		if part.IsPart {
			sum += part.Number
		}
	}
	// FIX: 1382231 too high
	fmt.Println("Part 1: ", sum)
}
