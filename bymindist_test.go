package constrain

import (
	"time"
	"testing"
	"github.com/franela/goblin"
	"github.com/TopoSimplify/opts"
)

func TestByMinDistRelation(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain by mindist relation", func() {
		g.It("should test constrain by context geometry", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 600 420, 580 440, 620 460, 620 500, 660 520, 720 500, 760 500, 760 440, 740 400, 700 440, 740 440 )")
			var cg_a = ctxGeom("POLYGON (( 660 360, 660 380, 680 380, 680 360, 660 360 ))")
			var cg_b = ctxGeom("POLYGON (( 660 440, 660 460, 680 460, 680 440, 660 440 ))")
			var cg_c = ctxGeom("POLYGON (( 660 540, 660 560, 680 560, 700 520, 660 540 ))")
			var hull = createNodes([][]int{{0, len(coords) - 1}}, coords)[0]
			var options = &opts.Opts{MinDist:10}

			g.Assert(ByMinDistRelation(options, hull, cg_a.AsContextGeometries())).IsTrue() // expands mindist
			g.Assert(ByMinDistRelation(options, hull, cg_b.AsContextGeometries())).IsFalse() //reduces mindist
			g.Assert(ByMinDistRelation(options, hull, cg_c.AsContextGeometries())).IsTrue() //expands mindist
		})
	})
}
