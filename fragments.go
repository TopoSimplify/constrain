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
	"simplex/common"
)

type scoreRelationFn func(float64) bool

const EpsilonDist = 1.0e-5

//constrain hulls at self intersection fragments - planar self-intersection
func atSelfIntersectFragments(
	hullDB *rtree.RTree,
	selfInters *ctx.ContextGeometries,
	atVertexSet *sset.SSet,
	scoreFn lnr.ScoreFn,
	scoreRelation scoreRelationFn,
) map[[2]int]*node.Node {
	var idxs []int
	var fragmentSize = 1
	var hsubs []*node.Node
	var hulls []*node.Node
	var unmerged = make(map[[2]int]*node.Node, 0)

	for _, inter := range selfInters.DataView() {
		if !inter.IsSelfVertex() {
			continue
		}
		hulls = common.NodesFromBoxes(
			knn.FindNeighbours(hullDB, inter, EpsilonDist),
		)

		idxs = common.AsInts(inter.Meta.SelfVertices.Values())
		for _, hull := range hulls {
			hsubs = split.AtIndex(hull, idxs, dp.NodeGeometry)

			if len(hsubs) == 0 && (hull.Range.Size() == fragmentSize) {
				hsubs = append(hsubs, hull)
			}

			if len(hsubs) == 0 {
				continue
			}

			hullDB.Remove(hull)
			keep, rm := merge.ContiguousFragmentsBySize(
				hsubs, hullDB, atVertexSet, unmerged, fragmentSize,
				scoreRelation, scoreFn, dp.NodeGeometry, EpsilonDist,
			)

			for _, h := range rm {
				hullDB.Remove(h)
			}

			for _, h := range keep {
				hullDB.Insert(h)
			}
		}
	}

	return unmerged
}
