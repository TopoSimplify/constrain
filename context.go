package constrain

import (
	"simplex/ctx"
	"simplex/lnr"
	"simplex/node"
	"simplex/relate"
	"simplex/deform"
	"github.com/intdxdt/rtree"
)

func ByGeometricRelation(self lnr.Linear, hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsGeomRelateValid(self, hull, cg)
}

func ByMinDistRelation(self lnr.Linear, hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDistRelateValid(self, hull, cg)
}

func BySideRelation(self lnr.Linear, hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDirRelateValid(self, hull, cg)
}

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func BySelfIntersection(self lnr.Linear, hull *node.Node, hulldb *rtree.RTree, selections *node.Nodes) bool {
	//assume hull is valid and proof otherwise
	var bln = true
	// find hull neighbours
	var hulls = deform.Select(self, hulldb, hull)
	for _, h := range hulls {
		//if bln & selection contains current hull : bln : false
		if bln && (h == hull) {
			bln = false //cmp &
		}
		selections.Push(h)
	}

	return bln
}
