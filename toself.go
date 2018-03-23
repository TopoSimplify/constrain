package constrain

import (
	"sort"
	"simplex/node"
	"simplex/pln"
	"simplex/opts"
	"simplex/ctx"
	"simplex/lnr"
	"simplex/common"
	"github.com/intdxdt/rtree"
)

//Constrain for planar self-intersection
func ToSelfIntersects(
	nodes []*node.Node, polyline *pln.Polyline, options *opts.Opts, constVerts []int,
) ([]*node.Node, bool, []int) {
	var atVertexSet = make(map[int]bool)
	if !options.KeepSelfIntersects {
		return nodes, true, []int{}
	}

	var hulldb = rtree.NewRTree(4)
	var selfInters = lnr.SelfIntersection(polyline)

	var data = make([]rtree.BoxObj, 0)
	for _, v := range nodes {
		data = append(data, v)
	}
	hulldb.Load(data)

	for _, inter := range selfInters.DataView() {
		if inter.IsNonPlanarVertex() {
			for _, v := range inter.Meta.NonPlanar {
				atVertexSet[v] = true
			}
		} else if inter.IsPlanarVertex() {
			for _, v := range inter.Meta.Planar {
				atVertexSet[v] = true
			}
		}
	}

	//update  const vertices if any
	//add const vertices as self inters
	for _, i := range constVerts {
		if atVertexSet[i] { //exclude already self intersects
			continue
		}
		atVertexSet[i] = true
		var pt = polyline.Coordinate(i)
		var cg = ctx.New(pt.Clone(), 0, -1).AsPlanarVertex()

		cg.Meta.Planar = append(cg.Meta.Planar, i)
		selfInters.Push(cg)
	}

	//constrain fragments around self intersects
	//try to merge fragments from first attempt
	splitAtSelfIntersects(hulldb, selfInters)

	nodes = common.NodesFromRtreeNodes(hulldb.All())
	sort.Sort(node.Nodes(nodes))

	indices := make([]int, 0, len(atVertexSet))
	for v := range atVertexSet {
		indices = append(indices, v)
	}
	sort.Ints(indices)
	return nodes, true, indices
}
