package graph

import (
	"fmt"
)

///// ///// ///// ///// ///// ///// GRAPH ///// ///// ///// ///// ///// /////

// Graph TODO:Write
type Graph struct {
	meta graphMeta
	vert map[int]vertex
	edge map[int]edge
}

// NewGraph TODO:Write
func NewGraph() *Graph {
	s := new(Graph)

	return s
}

func (g Graph) Test() {
	fmt.Println("Struct method test")
}

///// ///// ///// ///// ///// ///// GRAPH META ///// ///// ///// ///// ///// /////
type graphMeta struct {
	nodeID int
	maxID  int // Total number of worker nodes in complete graph
}

///// ///// ///// ///// ///// ///// VERTEX ///// ///// ///// ///// ///// /////
type vertex struct {
	id    int
	nType vertexType
}

type vertexType int

const (
	person vertexType = iota
	featPlace
	featJobTitle
)

///// ///// ///// ///// ///// ///// EDGE ///// ///// ///// ///// ///// /////
type edge struct {
	id int
}
