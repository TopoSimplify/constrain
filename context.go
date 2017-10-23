package constrain

import (
	"simplex/lnr"
	"simplex/knn"
	"simplex/node"
	"simplex/relate"
	"github.com/intdxdt/rtree"
)

//Constrain for context neighbours
// finds the collapsibility of hull with respect to context hull neighbours
// if hull is deformable, its added to selections
func ContextRelation(self lnr.Linear, ctxdb *rtree.RTree, hull *node.Node, selections *node.Nodes) bool {
	var bln = true
	var options = self.Options()
	// find context neighbours - if valid
	var ctxs = knn.FindNeighbours(ctxdb, hull, options.MinDist)
	for _, contxt := range ctxs {
		if !bln {
			break
		}

		cg := castAsContextGeom(contxt)
		if bln && options.GeomRelation {
			bln = relate.IsGeomRelateValid(self, hull, cg)
		}

		if bln && options.DistRelation {
			bln = relate.IsDistRelateValid(self, hull, cg)
		}

		if bln && options.DirRelation {
			bln = relate.IsDirRelateValid(self, hull, cg)
		}
	}

	if !bln {
		selections.Push(hull)
	}

	return bln
}
