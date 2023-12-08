package main

import (
	"fmt"
	"helpers"
	"strings"
)

type Nodes struct {
	Node map[string]Node
}

type Node struct {
	Name  string
	Left  string
	Right string
}

func NewNode(nodes string) Node {
	t := strings.Split(nodes, "=")
	name := strings.TrimSpace(t[0])

	tl := strings.Split(t[1], ",")[0]
	tl = strings.Replace(tl, "(", "", 1)
	leftNode := strings.TrimSpace(tl)

	tr := strings.Split(t[1], ",")[1]
	tr = strings.Replace(tr, ")", "", 1)
	rightNode := strings.TrimSpace(tr)

	return Node{
		Name:  name,
		Left:  leftNode,
		Right: rightNode,
	}
}

func MakeNodes(nodes []string) *Nodes {
	n := &Nodes{
		Node: make(map[string]Node),
	}
	for i := 2; i < len(nodes); i++ {
		t := strings.Split(nodes[i], "=")
		parent := strings.TrimSpace(t[0])

		n.Node[parent] = NewNode(nodes[i])
	}

	return n
}

func (n *Nodes) FindNode(seeker string) Node {
	return n.Node[seeker]
}

func main() {
	input := helpers.ReadFile("input.txt")

	ti := input[0]
	instructions := strings.Split(ti, "")

	nodes := MakeNodes(input)
	current := nodes.FindNode("AAA")

	steps := 0
	for i := 0; i < len(instructions); i++ {
		if current.Name == "ZZZ" {
			fmt.Printf("Found ZZZ in {%d} steps!\n", steps)
			break
		}

		if instructions[i] == "L" {
			current = nodes.FindNode(current.Left)
		}

		if instructions[i] == "R" {
			current = nodes.FindNode(current.Right)
		}

		steps += 1
		if i == len(instructions)-1 {
			i = -1
		}
	}
}
