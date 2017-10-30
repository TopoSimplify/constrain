package constrain

import (
	"simplex/node"
	"simplex/deform"
	"github.com/intdxdt/rtree"
	"simplex/opts"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func ByFeatureClassIntersection(options *opts.Opts, hull *node.Node, hulldb *rtree.RTree, selections *node.Nodes) bool {
	var bln = true
	//find hull neighbours
	var hulls = deform.SelectFeatureClass(options, hulldb, hull)
	for _, h := range hulls {
		//if bln & selection contains current hull : bln : false
		if bln && (h == hull) {
			bln = false // cmp ref
		}
		selections.Push(h)
	}
	return bln
}
