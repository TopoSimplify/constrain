package constrain

import (
	"simplex/dp"
	"simplex/knn"
	"simplex/ctx"
	"simplex/node"
	"simplex/split"
	"simplex/common"
	"github.com/intdxdt/rtree"
)

const EpsilonDist = 1.0e-5

//constrain hulls at self intersection fragments - planar self-intersection
func splitAtSelfIntersects(hullDB *rtree.RTree, selfInters *ctx.ContextGeometries) {
	var tokens []*node.Node
	var hulls []*node.Node

	for _, inter := range selfInters.DataView() {
		var idxs []int
		if inter.IsPlanarVertex() {
			idxs = inter.Meta.Planar
		} else if inter.IsNonPlanarVertex() {
			idxs = inter.Meta.NonPlanar
		}
		if len(idxs) == 0 {
			continue
		}

		hulls = common.NodesFromBoxes(
			knn.FindNeighbours(hullDB, inter, EpsilonDist),
		)

		for _, hull := range hulls {
			tokens = split.AtIndex(hull, idxs, dp.NodeGeometry)

			if len(tokens) == 0 {
				continue
			}

			hullDB.Remove(hull)
			for _, h := range tokens {
				hullDB.Insert(h)
			}
		}
	}
}
