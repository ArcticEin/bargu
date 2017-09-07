package bargu

import "bargu/graph"
import "fmt"

type Node struct {
	st *graph.Graph
}

func NewNode() *Node {
	n := new(Node)
	n.st = graph.NewGraph()
	return n
}

func (n Node) Test() {
	fmt.Println("Node test")
	n.st.Test()
}
