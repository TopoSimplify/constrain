package constrain

import (
	"simplex/node"
	"simplex/ctx"
	"simplex/relate"
)

func ByGeometricRelation(hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsGeomRelateValid(hull, cg)
}
