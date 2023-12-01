package main

import (
	"fmt"
	"strings"
	"util"
)

type coordinate struct {
	x int
	y int
}

func (c *coordinate) moveLeft() {
	c.x -= 1
}

func (c *coordinate) moveRight() {
	c.x += 1
}

func (c *coordinate) moveUp() {
	c.y += 1
}

func (c *coordinate) moveDown() {
	c.y -= 1
}

func (c *coordinate) moveUpLeft() {
	c.moveUp()
	c.moveLeft()
}

func (c *coordinate) moveUpRight() {
	c.moveUp()
	c.moveRight()
}

func (c *coordinate) moveDownLeft() {
	c.moveDown()
	c.moveLeft()
}

func (c *coordinate) moveDownRight() {
	c.moveDown()
	c.moveRight()
}

func (c *coordinate) follow(pos *coordinate) {
	xPos := c.x - pos.x
	yPos := c.y - pos.y

	if xPos == 1 && yPos == 2 {
		c.moveDownLeft()
		return
	}

	if xPos == -2 && yPos == 1 {
		c.moveUpLeft()
		return
	}

	if xPos == 1 && yPos == -2 {
		c.moveDownRight()
		return
	}

	if xPos == -1 && yPos == 2 {
		c.moveUpLeft()
		return
	}

	if xPos == -1 && yPos == -2 {
		c.moveUpRight()
		return
	}

	if xPos == -2 {
		c.moveRight()
		return
	}

	if xPos == 2 {
		c.moveLeft()
		return
	}

	if yPos == -2 {
		c.moveUp()
		return
	}

	if yPos == 2 {
		c.moveDown()
		return
	}
}

type visited struct {
	coordinates []coordinate
}

func hasVisited() {}

func part1() {
	moves := util.ReadFile("input2.txt")

	head := &coordinate{
		x: 0,
		y: 0,
	}

	tail := &coordinate{
		x: 0,
		y: 0,
	}

	for i := 0; i < len(moves); i++ {
		instruction := strings.Fields(moves[i])

		if strings.Contains(moves[i], "R") {
			for j := 0; j < util.Stoi(instruction[1]); j++ {
				head.moveRight()
				tail.follow(head)
				fmt.Println(tail, head)
			}
		}

		if strings.Contains(moves[i], "D") {
			for j := 0; j < util.Stoi(instruction[1]); j++ {
				head.moveDown()
				tail.follow(head)
				fmt.Println(tail, head)
			}
		}

		if strings.Contains(moves[i], "L") {
			for j := 0; j < util.Stoi(instruction[1]); j++ {
				head.moveLeft()
				tail.follow(head)
				fmt.Println(tail, head)
			}
		}

		if strings.Contains(moves[i], "U") {
			for j := 0; j < util.Stoi(instruction[1]); j++ {
				head.moveUp()
				tail.follow(head)
				fmt.Println(tail, head)
			}
		}
	}
}

func main() {
	part1()
}
