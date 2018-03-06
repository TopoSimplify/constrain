package constrain

import (
	"simplex/node"
	"simplex/ctx"
	"simplex/relate"
)

func ByGeometricRelation(hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsGeomRelateValid(hull, cg)
}
