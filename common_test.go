package constrain

import (
    "fmt"
    "simplex/dp"
    "simplex/pln"
    "simplex/rng"
    "simplex/node"
    "github.com/intdxdt/geom"
    "simplex/ctx"
    "simplex/db"
)

func DebugPrintNodes(ns []*node.Node) {
    for _, n := range ns {
        fmt.Println(n.Geom.WKT())
    }
}

func ctxGeom(wkt string) *ctx.ContextGeometry {
    return ctx.New(geom.NewGeometry(wkt), 0, -1)
}

func linearCoords(wkt string) []*geom.Point {
    return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createNodes(indxs [][]int, coords []*geom.Point) []*node.Node {
    poly := pln.New(coords)
    hulls := make([]*node.Node, 0)
    for _, o := range indxs {
        hulls = append(hulls, node.NewFromPolyline(poly, rng.NewRange(o[0], o[1]), dp.NodeGeometry))
    }
    return hulls
}

//hull db
func hullsDB(ns []*node.Node) *db.DB {
    database := db.NewDB(8)
    for _, n := range ns {
        database.Insert(n)
    }
    return database
}

//hull geom
func hullGeom(coords []*geom.Point) geom.Geometry {
    var g geom.Geometry

    if len(coords) > 2 {
        g = geom.NewPolygon(coords)
    } else if len(coords) == 2 {
        g = geom.NewLineString(coords)
    } else {
        g = coords[0].Clone()
    }
    return g
}
