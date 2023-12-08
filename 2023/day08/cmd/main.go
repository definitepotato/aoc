package main

import (
	"fmt"
	"helpers"
	"log"
	"strings"
)

type Nodes struct {
	Nodes []*Node
}

type Node struct {
	Left  string
	Right string
	Name  string
}

func NewNode(nodes string) *Node {
	t := strings.Split(nodes, "=")
	mainNode := strings.TrimSpace(t[0])

	tl := strings.Split(t[1], ",")[0]
	tl = strings.Replace(tl, "(", "", 1)
	leftNode := strings.TrimSpace(tl)

	tr := strings.Split(t[1], ",")[1]
	tr = strings.Replace(tr, ")", "", 1)
	rightNode := strings.TrimSpace(tr)

	return &Node{
		Name:  mainNode,
		Left:  leftNode,
		Right: rightNode,
	}
}

func MakeNodes(nodes []string) *Nodes {
	n := &Nodes{}
	for i := 2; i < len(nodes); i++ {
		n.Nodes = append(n.Nodes, NewNode(nodes[i]))
	}

	return n
}

func (n *Nodes) FindNode(seeker string) (*Node, error) {
	for _, node := range n.Nodes {
		if node.Name == seeker {
			return node, nil
		}
	}

	return nil, fmt.Errorf("node not found: {%s}", seeker)
}

func main() {
	input := helpers.ReadFile("input.txt")

	ti := input[0]
	instructions := strings.Split(ti, "")

	nodes := MakeNodes(input)
	current := nodes.Nodes[0]

	steps := 0
	for i := 0; i < len(instructions); i++ {
		if current.Name == "ZZZ" {
			fmt.Printf("Found ZZZ! It took {%d} steps\n", steps)
			break
		}

		if instructions[i] == "L" {
			next, err := nodes.FindNode(current.Left)
			if err != nil {
				log.Fatal(err)
			}

			current = next
		}

		if instructions[i] == "R" {
			next, err := nodes.FindNode(current.Right)
			if err != nil {
				log.Fatal(err)
			}

			current = next
		}

		steps += 1
		if i == len(instructions)-1 {
			i = -1
		}
	}
}
