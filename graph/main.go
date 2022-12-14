package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Graph[T any] struct {
	nodes []*Node[T]
}

type Node[T any] struct {
	value T
	edges []*Node[T]
}

func NewGraph[T any]() *Graph[T] {
	return &Graph[T]{}
}

func (g *Graph[T]) AddNode(val T) *Node[T] {
	node := &Node[T]{value: val}
	g.nodes = append(g.nodes, node)
	return node
}

func (g *Graph[T]) AddEdge(from, to *Node[T]) {
	from.edges = append(from.edges, to)
}

func (g *Graph[T]) HasNode(node *Node[T]) bool {
	for _, n := range g.nodes {
		if n == node {
			return true
		}
	}

	return false
}

func (g *Graph[T]) HasEdge(from, to *Node[T]) bool {
	for _, n := range from.edges {
		if n == to {
			return true
		}
	}

	return false
}

func (g *Graph[T]) Print() {
	data := [][]string{}
	for _, node := range g.nodes {
		var edges string
		for _, edge := range node.edges {
			edges += fmt.Sprintf("%v ", edge.value)
		}

		data = append(data, []string{fmt.Sprintf("%v", node.value), strings.TrimSpace(edges)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Node", "Edges"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

func main() {
	g := NewGraph[int]()
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)

	g.AddEdge(n1, n3)
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n4)

	g.Print()
}
