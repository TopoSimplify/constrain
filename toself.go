package constrain

import (
	"simplex/ctx"
	"simplex/lnr"
	"simplex/node"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/rtree"
)

//Constrain for planar self-intersection
func ToSelfIntersects(self lnr.Linear, constVerts []int, scoreRelation scoreRelationFn) (*deque.Deque, bool, *sset.SSet) {
	var atVertexSet *sset.SSet
	var polyline = self.Polyline()
	var options = self.Options()
	if !options.KeepSelfIntersects {
		return self.NodeQueue(), true, atVertexSet
	}

	var hulldb = rtree.NewRTree(16)
	var selfInters = lnr.SelfIntersection(self)

	var data = make([]rtree.BoxObj, 0)
	for _, v := range *self.NodeQueue().DataView() {
		data = append(data, v.(*node.Node))
	}
	hulldb.Load(data)

	atVertexSet = sset.NewSSet(cmp.Int)
	for _, inter := range selfInters.DataView() {
		if inter.IsSelfVertex() {
			atVertexSet = atVertexSet.Union(inter.Meta.SelfVertices)
		}
	}

	//update  const vertices if any
	//add const vertices as self inters
	for _, i := range constVerts {
		if atVertexSet.Contains(i) { //exclude already self intersects
			continue
		}
		atVertexSet.Add(i)
		var pt = polyline.Coordinate(i)
		var cg = ctx.New(pt.Clone(), i, i).AsSelfVertex()

		cg.Meta.SelfVertices = sset.NewSSet(cmp.Int, 4).Add(i)
		cg.Meta.SelfNonVertices = sset.NewSSet(cmp.Int, 4)
		selfInters.Push(cg)
	}

	//constrain fragments around self intersects
	//try to merge fragments from first attempt
	var mcount = 2
	for mcount > 0 {
		fragments := atSelfIntersectFragments(self, hulldb, selfInters, atVertexSet, scoreRelation)
		if len(fragments) == 0 {
			break
		}
		mcount += -1
	}
	return nodesFromRtreeNodes(hulldb.All()).Sort().AsDeque(), true, atVertexSet
}
