package constrain

import (
	"github.com/franela/goblin"
	"time"
	"simplex/pln"
	"simplex/opts"
	"github.com/intdxdt/deque"
	"simplex/offset"
	"simplex/node"
	"testing"
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
			var queue = deque.NewDeque()
			for _, n := range nodes {
				queue.Append(n)
			}

			var constVerts = []int{10}
			var scoreFn = offset.MaxOffset
			var scoreRelationFn = func(f float64) bool {
				if f <= options.Threshold {
					return true
				}
				return false
			}

			options.KeepSelfIntersects = false
			var que, bln, set = ToSelfIntersects(
				queue, polyline, options, constVerts, scoreFn, scoreRelationFn,
			)
			g.Assert(bln).IsTrue()

			options.KeepSelfIntersects = true
			que, bln, set = ToSelfIntersects(
				queue, polyline, options,
				constVerts, scoreFn, scoreRelationFn,
			)

			g.Assert(bln).IsTrue()
			g.Assert(bln).IsTrue()
			nodes = []*node.Node{}
			for _, n := range *que.DataView() {
				nodes = append(nodes, n.(*node.Node))
			}
			g.Assert(set.Values()).Equal([]interface{}{3, 7, 10})
			g.Assert(len(nodes)).Equal(6)

			//que, bln, set = ToSelfIntersects(
			//	queue, polyline, options,
			//	constVerts, scoreFn, scoreRelationFn,
			//)

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
			var queue = deque.NewDeque()
			for _, n := range nodes {
				queue.Append(n)
			}

			var constVerts = []int{10, 7}
			var scoreFn = offset.MaxOffset
			var scoreRelationFn = func(f float64) bool {
				if f <= options.Threshold {
					return true
				}
				return false
			}

			options.KeepSelfIntersects = false
			var que, bln, set = ToSelfIntersects(
				queue, polyline, options, constVerts, scoreFn, scoreRelationFn,
			)
			g.Assert(bln).IsTrue()

			options.KeepSelfIntersects = true
			que, bln, set = ToSelfIntersects(
				queue, polyline, options,
				constVerts, scoreFn, scoreRelationFn,
			)
			//DebugPrintNodes(nodes)
			//for _, h := range *que.DataView(){
			//	fmt.Println(h.(*node.Node).Geom.WKT())
			//}
			g.Assert(que.Len()).Equal(5)
			g.Assert(set.Values()).Equal([]interface{}{3, 7, 10})
		})
	})
}
