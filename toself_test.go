package constrain

import (
	"time"
	"fmt"
	"testing"
	"simplex/pln"
	"simplex/opts"
	"github.com/franela/goblin"
	"strings"
	"simplex/node"
)

func TestToSelfIntersects(t *testing.T) {
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

			g.Assert(len(nodes)).Equal(3)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{}

			options.KeepSelfIntersects = false
			var que, bln, set = ToSelfIntersects(
				queue, polyline, options, constVerts,
			)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(3)
			g.Assert(len(set)).Equal(0)

			constVerts = []int{10}
			options.KeepSelfIntersects = true
			que, bln, set = ToSelfIntersects(
				queue, polyline, options, constVerts,
			)

			g.Assert(bln).IsTrue()
			g.Assert(set).Equal([]int{0, 1, 3, 7, 9, 10})
		})


		g.It("should test with cont  verts", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 780 480, 750 470, 760 500, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
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

			fmt.Println(strings.Repeat("--", 80))

			g.Assert(len(nodes)).Equal(3)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{10}

			options.KeepSelfIntersects = false
			var que, bln, set = ToSelfIntersects(
				queue, polyline, options, constVerts,
			)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(3)
			g.Assert(len(set)).Equal(0)

			constVerts = []int{10}
			options.KeepSelfIntersects = true
			que, bln, set = ToSelfIntersects(
				queue, polyline, options, constVerts,
			)

			g.Assert(bln).IsTrue()
			nodes = []*node.Node{}
			for _, n := range que {
				fmt.Println(n.Geom.WKT())
			}
			g.Assert(len(que)).Equal(6)
			g.Assert(set).Equal([]int{3, 7,  10})
		})

		g.It("should test constrain to self intersects - merge fragments", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 740 380, 720 440, 760 460, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			//var cong = geom.NewPolygonFromWKT("POLYGON (( 780 560, 780 580, 800 580, 800 560, 780 560 ))")
			var polyline = pln.New(coords)
			options := &opts.Opts{
				Threshold:              300.0,
				MinDist:                300.0,
				RelaxDist:              300.0,
				KeepSelfIntersects:     true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var nodes = createNodes([][]int{{0, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 9}, {9, 11}}, coords)
			g.Assert(len(nodes)).Equal(7)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{10}

			options.KeepSelfIntersects = false
			var que, bln, _ = ToSelfIntersects(
				queue, polyline, options, constVerts,
			)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(len(nodes))

			options.KeepSelfIntersects = true
			que, bln, _ = ToSelfIntersects(
				queue, polyline, options, constVerts,
			)
			g.Assert(len(que)).Equal(9)
		})
	})
}
