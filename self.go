package constrain

import (
	"simplex/lnr"
	"simplex/node"
	"simplex/deform"
	"github.com/intdxdt/rtree"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func SelfIntersection(self lnr.Linear, hull *node.Node, hulldb *rtree.RTree, selections *node.Nodes) bool {
	//assume hull is valid and proof otherwise
	var bln = true
	// find hull neighbours
	hulls := deform.Select(self, hulldb, hull)
	for _, h := range hulls {
		//if bln & selection contains current hull : bln : false
		if bln && (h == hull) {
			bln = false //cmp &
		}
		selections.Push(h)
	}

	return bln
}
