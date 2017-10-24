package constrain

import (
	"simplex/node"
	"github.com/intdxdt/rtree"
)

type scoreRelationFn func(float64) bool

const EpsilonDist = 1.0e-5

//Convert slice of interface to ints
func asInts(iter []interface{}) []int {
	var ints = make([]int, len(iter))
	for i, o := range iter {
		ints[i] = o.(int)
	}
	return ints
}

//node.Nodes from Rtree boxes
func nodesFromBoxes(iter []rtree.BoxObj) *node.Nodes {
	var self = node.NewNodes(len(iter))
	for _, h := range iter {
		self.Push(h.(*node.Node))
	}
	return self
}

//node.Nodes from Rtree nodes
func nodesFromRtreeNodes(iter []*rtree.Node) *node.Nodes {
	var self = node.NewNodes(len(iter))
	for _, h := range iter {
		self.Push(h.GetItem().(*node.Node))
	}
	return self
}
