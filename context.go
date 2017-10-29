package constrain

import (
	"simplex/ctx"
	"simplex/node"
	"simplex/relate"
	"simplex/deform"
	"github.com/intdxdt/rtree"
	"simplex/opts"
)

func ByGeometricRelation(hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsGeomRelateValid(hull, cg)
}

func ByMinDistRelation(options *opts.Opts, hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDistRelateValid(options, hull, cg)
}

func BySideRelation(hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDirRelateValid(hull, cg)
}

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
