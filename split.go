package constrain

import (
	"github.com/TopoSimplify/dp"
	"github.com/TopoSimplify/knn"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/split"
	"github.com/intdxdt/rtree"
)

const EpsilonDist = 1.0e-5

//constrain hulls at self intersection fragments - planar self-intersection
func splitAtSelfIntersects(hullDB *rtree.RTree, selfInters *ctx.ContextGeometries) {
	var tokens []*node.Node
	var objects []*rtree.Obj

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

		objects = knn.FindNeighbours(hullDB, inter.Geom, EpsilonDist)
		for i  := range objects {
			var obj = objects[i]
			var hull = obj.Object.(*node.Node)
			tokens = split.AtIndex(hull, idxs, dp.NodeGeometry)

			if len(tokens) == 0 {
				continue
			}

			hullDB.RemoveObj(obj)
			for _, h := range tokens {
				hullDB.Insert(rtree.Object(obj.Id, h.Bounds(), h))
			}
		}
	}
}
