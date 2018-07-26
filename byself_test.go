package constrain

import (
	"time"
	"testing"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/node"
	"github.com/franela/goblin"
	"github.com/TopoSimplify/hdb"
)

func TestBySelfIntersection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain by mindist relation", func() {
		g.It("should test constrain by context geometry", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 600 420, 580 440, 620 460, 620 500, 660 520, 720 500, 760 500, 760 440, 740 400, 700 440, 740 440 )")
			var hulls = createNodes([][]int{{0, 4}, {4, len(coords) - 1}}, coords)
			var db = hdb.NewHdb().Load(hulls)
			var sels = []*node.Node{}
			var options = &opts.Opts{MinDist: 10}
			g.Assert(BySelfIntersection(options, hulls[0], db, &sels)).IsTrue()
			g.Assert(BySelfIntersection(options, hulls[1], db, &sels)).IsTrue()

			coords = linearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
			hulls = createNodes([][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, { 17, len(coords)-1}}, coords)
			db = hdb.NewHdb().Load(hulls)
			g.Assert(BySelfIntersection(options, hulls[0], db, &sels)).IsTrue()
			g.Assert(BySelfIntersection(options, hulls[1], db, &sels)).IsFalse()
			g.Assert(BySelfIntersection(options, hulls[2], db, &sels)).IsTrue()
			g.Assert(BySelfIntersection(options, hulls[3], db, &sels)).IsTrue()
			g.Assert(BySelfIntersection(options, hulls[4], db, &sels)).IsTrue()

		})
	})
}
