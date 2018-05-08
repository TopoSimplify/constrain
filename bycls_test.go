package constrain

import (
	"time"
	"testing"
	"github.com/TopoSimplify/dp"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/offset"
	"github.com/franela/goblin"
)

func TestByFeatureClassIntersection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain by mindist relation", func() {
		g.It("should test constrain by context geometry", func() {
			g.Timeout(1 * time.Hour)

			var options = &opts.Opts{MinDist: 10}
			var coords = linearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
			var hulls = createNodes([][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, len(coords) - 1}}, coords)
			var inst = dp.New(coords, options, offset.MaxOffset)

			for _, h := range hulls {
				h.Instance = inst
			}

			var db = hullsDB(hulls)
			var sels = []*node.Node{}

			coords = linearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			for _, h := range hulls {
				h.Instance = inst
				db.Insert(h)
			}
			var q1 = hulls[0]
			coords = linearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			for _, h := range hulls {
				h.Instance = inst
				db.Insert(h)
			}
			var q2 = hulls[0]

			g.Assert(ByFeatureClassIntersection(options, q1, db, &sels)).IsFalse()
			g.Assert(ByFeatureClassIntersection(options, q2, db, &sels)).IsTrue()

		})
	})
}
