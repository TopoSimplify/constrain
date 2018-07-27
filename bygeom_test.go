package constrain

import (
	"time"
	"testing"
	"github.com/franela/goblin"
)

func TestByGeometricRelation(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain by geometric relation", func() {
		g.It("should test constrain by context geometry", func() {
			g.Timeout(1 * time.Hour)
			var coords = linearCoords("LINESTRING ( 600 420, 580 440, 620 460, 620 500, 660 520, 720 500, 760 500, 760 440, 740 400, 700 440, 740 440 )")
			var cg_a = ctxGeom("POLYGON (( 660 360, 660 380, 680 380, 680 360, 660 360 ))")
			var cg_b = ctxGeom("POLYGON (( 660 420, 660 460, 680 460, 680 420, 660 420 ))")
			var cg_c = ctxGeom("POLYGON (( 660 500, 660 540, 680 540, 680 500, 660 500 ))")
			var hull = createNodes([][]int{{0, len(coords) - 1}}, coords)[0]
			g.Assert(ByGeometricRelation(&hull, cg_a.AsContextGeometries())).IsTrue() //disjoint
			g.Assert(ByGeometricRelation(&hull, cg_b.AsContextGeometries())).IsFalse() //disjoint
			g.Assert(ByGeometricRelation(&hull, cg_c.AsContextGeometries())).IsFalse() //intersects
		})
	})
}
