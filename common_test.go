package constrain

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/dp"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/ctx"
	"github.com/intdxdt/iter"
)

func ctxGeom(wkt string) *ctx.ContextGeometry {
	return ctx.New(geom.ReadGeometry(wkt), 0, -1)
}

func linearCoords(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createNodes(id *iter.Igen, indxs [][]int, coords []geom.Point) []node.Node {
	var poly  = pln.New(coords)
	var hulls = make([]node.Node, 0)
	for i := range indxs {
		r := rng.Range(indxs[i][0], indxs[i][1])
		hulls = append(hulls, node.CreateNode(id, poly.SubCoordinates(r), r, dp.NodeGeometry))
	}
	return hulls
}

