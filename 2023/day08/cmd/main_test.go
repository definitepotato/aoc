package main

import (
	"helpers"
	"testing"
)

func TestNewNode(t *testing.T) {
	t.Run("Create new node", func(t *testing.T) {
		have := "AAA = (BBB, CCC)"
		got := NewNode(have)

		if got.Name != "AAA" {
			t.Errorf("got {%s}, wanted AAA", got.Name)
		}

		if got.Left != "BBB" {
			t.Errorf("got {%s}, wanted BBB", got.Left)
		}

		if got.Right != "CCC" {
			t.Errorf("got {%s}, wanted CCC", got.Right)
		}
	})
}

func TestMakeNodes(t *testing.T) {
	t.Run("Making new nodes from input", func(t *testing.T) {
		input := helpers.ReadFile("../test.txt")
		got := MakeNodes(input)
		want := len(got.Node)

		if want != 3 {
			t.Errorf("got {%d}, wanted 3", want)
		}
	})
}

func TestGetFirstNodeName(t *testing.T) {
	t.Run("Find the initial node starting point", func(t *testing.T) {
		have := helpers.ReadFile("../test.txt")
		got := GetFirstNodeName(have)
		want := "AAA"

		if got != want {
			t.Errorf("got {%s}, wanted {%s}", got, want)
		}
	})
}

func TestFindNode(t *testing.T) {
	t.Run("Find node in nodes", func(t *testing.T) {
		input := helpers.ReadFile("../test.txt")
		got := MakeNodes(input)
		want := got.FindNode("ZZZ")

		if want.Name != "ZZZ" {
			t.Errorf("got {%s}, wanted ZZZ", want.Name)
		}
	})
}
