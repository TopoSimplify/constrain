package constrain

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/ctx"
)

func ctxGeom(wkt string) *ctx.ContextGeometry {
	return ctx.New(geom.ReadGeometry(wkt), 0, -1)
}
