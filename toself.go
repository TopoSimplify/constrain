package constrain

import (
	"sort"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/hdb"
	"github.com/intdxdt/iter"
)

//Constrain for planar self-intersection
func ToSelfIntersects( id *iter.Igen,
	nodes []node.Node, polyline pln.Polyline, options *opts.Opts, constVerts []int,
) ([]node.Node, bool, []int) {
	var atVertexSet = make(map[int]bool)
	if !options.PlanarSelf {
		return nodes, true, []int{}
	}

	var hulldb = hdb.NewHdb().Load(nodes)
	var planar, nonPlanar = options.PlanarSelf, options.NonPlanarSelf
	var selfInters = lnr.SelfIntersection(polyline, planar, nonPlanar)

	for _, inter := range selfInters.DataView() {
		var indices = inter.Meta.Planar
		if inter.IsNonPlanarVertex() {
			indices = inter.Meta.NonPlanar
		}
		for _, v := range indices {
			atVertexSet[v] = true
		}
	}

	//update  const vertices if any
	//add const vertices as self inters
	for _, i := range constVerts {
		if atVertexSet[i] { //exclude already self intersects
			continue
		}
		atVertexSet[i] = true
		var pt = polyline.Coordinates.Pt(i)
		var cg = ctx.New(pt, 0, -1).AsPlanarVertex()

		cg.Meta.Planar = append(cg.Meta.Planar, i)
		selfInters.Push(cg)
	}

	splitAtSelfIntersects(id, hulldb, selfInters)

	nodes = make([]node.Node, 0, len(nodes))
	for _,n := range hulldb.All() {
		nodes = append(nodes, *n)
	}

	sort.Sort(node.Nodes(nodes))

	var indices = make([]int, 0, len(atVertexSet))
	for v := range atVertexSet {
		indices = append(indices, v)
	}
	sort.Ints(indices)
	return nodes, true, indices
}
