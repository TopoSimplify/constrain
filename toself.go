package constrain

import (
	"simplex/ctx"
	"simplex/pln"
	"simplex/lnr"
	"simplex/node"
	"simplex/opts"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/rtree"
	"simplex/common"
	"sort"
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
			for _, v := range inter.Meta.NonPlanarVertices.Values() {
				atVertexSet[v.(int)] = true
			}
		} else if inter.IsPlanarVertex() {
			for _, v := range inter.Meta.PlanarVertices.Values() {
				atVertexSet[v.(int)] = true
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

		cg.Meta.PlanarVertices = sset.NewSSet(cmp.Int, 4).Add(i)
		cg.Meta.NonPlanarVertices = sset.NewSSet(cmp.Int, 4)
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
