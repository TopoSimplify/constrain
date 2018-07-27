package constrain

import (
	"github.com/franela/goblin"
	"time"
	"testing"
)

func TestBySideRelation(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain by side relation", func() {
		g.It("should test constrain by context geometry", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 600 420, 580 440, 620 460, 620 500, 660 520, 720 500, 760 500, 760 440, 740 400, 700 440, 740 440 )")
			var cgA = ctxGeom("POLYGON (( 660 360, 660 380, 680 380, 680 360, 660 360 ))")
			var cgB = ctxGeom("POLYGON (( 660 440, 660 460, 680 460, 680 440, 660 440 ))")
			var cgC = ctxGeom("POLYGON (( 660 540, 660 560, 680 560, 700 520, 660 540 ))")
			var hull = createNodes([][]int{{0, len(coords) - 1}}, coords)[0]

			g.Assert(BySideRelation(&hull, cgA.AsContextGeometries())).IsTrue()
			g.Assert(BySideRelation(&hull, cgB.AsContextGeometries())).IsFalse()
			g.Assert(BySideRelation(&hull, cgC.AsContextGeometries())).IsTrue()
		})
	})
}
