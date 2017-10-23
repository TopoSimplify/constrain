package constrain

import (
	"simplex/lnr"
	"simplex/node"
	"simplex/deform"
	"github.com/intdxdt/rtree"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func FeatureClassIntersection(self lnr.Linear, hull *node.Node, hulldb *rtree.RTree, selections *node.Nodes) bool {
	var bln = true
	//find hull neighbours
	var hulls = deform.SelectFeatureClass(self, hulldb, hull)
	for _, h := range hulls {
		//if bln & selection contains current hull : bln : false
		if bln && (h == hull) {
			bln = false // cmp ref
		}
		selections.Push(h)
	}
	return bln
}
