package constrain

import (
	"testing"
	"simplex/dp"
	"simplex/pln"
	"simplex/rng"
	"simplex/node"
	"simplex/opts"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
	"github.com/intdxdt/deque"
	"simplex/offset"
	"fmt"
	"time"
)

func DebugPrintNodes(ns []*node.Node) {
	for _, n := range ns {
		fmt.Println(n.Geom.WKT())
	}
}

func TestCommon(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain", func() {
		g.It("should test constrain to self intersects", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 740 380, 720 440, 760 460, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			//var cong = geom.NewPolygonFromWKT("POLYGON (( 780 560, 780 580, 800 580, 800 560, 780 560 ))")
			var polyline = pln.New(coords)
			options := &opts.Opts{
				Threshold:              1.0,
				MinDist:                1.0,
				RelaxDist:              1.0,
				KeepSelfIntersects:     true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var nodes = createNodes([][]int{{0, 5}, {5, 9}, {9, 11}}, coords)
			//DebugPrintNodes(nodes)
			g.Assert(len(nodes)).Equal(3)
			var queue = deque.NewDeque()
			for _, n := range nodes {
				queue.Append(n)
			}

			var constVerts = []int{3}
			var scoreFn = offset.MaxOffset
			var scoreRelationFn = func(f float64) bool {
				if f <= options.Threshold {
					return true
				}
				return false
			}
			var que, bln, set = ToSelfIntersects(
				queue, polyline, options,
				constVerts, scoreFn, scoreRelationFn,
			)
			g.Assert(bln).IsTrue()
			nodes = []*node.Node{}
			for _, n := range *que.DataView() {
				nodes = append(nodes, n.(*node.Node))
			}
			g.Assert(set.Values()).Equal([]interface{}{3, 7})
			g.Assert(len(nodes)).Equal(5)

			//que, bln, set = ToSelfIntersects(
			//	queue, polyline, options,
			//	constVerts, scoreFn, scoreRelationFn,
			//)

		})
	})
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
