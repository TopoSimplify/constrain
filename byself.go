package constrain

import (
	"simplex/node"
	"simplex/deform"
	"simplex/opts"
	"github.com/intdxdt/rtree"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func BySelfIntersection(options *opts.Opts, hull *node.Node, hulldb *rtree.RTree, selections *node.Nodes) bool {
	//assume hull is valid and proof otherwise
	var bln = true
	// find hull neighbours
	var hulls = deform.Select(options, hulldb, hull)
	for _, h := range hulls {
		//if bln & selection contains current hull : bln : false
		if bln && (h == hull) {
			bln = false //cmp &
		}
		selections.Push(h)
	}

	return bln
}
