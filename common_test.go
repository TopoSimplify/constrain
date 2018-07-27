package constrain

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/dp"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/ctx"
)

func ctxGeom(wkt string) *ctx.ContextGeometry {
	return ctx.New(geom.ReadGeometry(wkt), 0, -1)
}

func linearCoords(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createNodes(indxs [][]int, coords []geom.Point) []node.Node {
	poly := pln.New(coords)
	hulls := make([]node.Node, 0)
	for _, o := range indxs {
		r := rng.Range(o[0], o[1])
		hulls = append(hulls, node.CreateNode(poly.SubCoordinates(r), r, dp.NodeGeometry))
	}
	return hulls
}

