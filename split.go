package constrain

import (
	"github.com/TopoSimplify/common"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/hdb"
	"github.com/TopoSimplify/knn"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/split"
	"github.com/intdxdt/iter"
)

const EpsilonDist = 1.0e-5

//constrain hulls at self intersection fragments - planar self-intersection
func splitAtSelfIntersects(id *iter.Igen, db *hdb.Hdb, selfInters *ctx.ContextGeometries) {
	var tokens []node.Node
	var nodes []*node.Node
	var hull *node.Node

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

		nodes = knn.ContextNeighbours(db, inter.Geom, EpsilonDist)
		for i := range nodes {
			hull = nodes[i]
			tokens = split.AtIndex(id, hull, idxs, common.Geometry)
			if len(tokens) == 0 {
				continue
			}
			db.Remove(hull).Load(tokens)
		}
	}
}
