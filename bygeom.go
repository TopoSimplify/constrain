package constrain

import (
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/ctx"
	"simplex/relate"
)

func ByGeometricRelation(hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsGeomRelateValid(hull, cg)
}
