package constrain

import (
	"simplex/lnr"
	"simplex/dp"
	"simplex/knn"
	"simplex/ctx"
	"simplex/node"
	"simplex/split"
	"simplex/merge"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/rtree"
)

//constrain hulls at self intersection fragments - planar self-intersection
func atSelfIntersectFragments(self lnr.Linear, hulldb *rtree.RTree,
	selfInters *ctx.ContextGeometries, atVertexSet *sset.SSet, scoreRelation scoreRelationFn) map[[2]int]*node.Node {
	var fragmentSize = 1
	var hsubs []*node.Node
	var hulls *node.Nodes
	var idxs []int
	var unmerged = make(map[[2]int]*node.Node, 0)

	for _, inter := range selfInters.DataView() {
		if !inter.IsSelfVertex() {
			continue
		}

		hulls = nodesFromBoxes(knn.FindNeighbours(hulldb, inter, EpsilonDist)).Sort()

		idxs = asInts(inter.Meta.SelfVertices.Values())
		for _, hull := range hulls.DataView() {
			hsubs = split.AtIndex(hull, idxs, dp.NodeGeometry)

			if len(hsubs) == 0 && (hull.Range.Size() == fragmentSize) {
				hsubs = append(hsubs, hull)
			}

			if len(hsubs) == 0 {
				continue
			}

			hulldb.Remove(hull)
			keep, rm := merge.ContiguousFragmentsBySize(self,
				hsubs, hulldb, atVertexSet, unmerged, fragmentSize,
				scoreRelation, dp.NodeGeometry, EpsilonDist,
			)

			for _, h := range rm {
				hulldb.Remove(h)
			}

			for _, h := range keep {
				hulldb.Insert(h)
			}
		}
	}

	return unmerged
}
