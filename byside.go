package constrain

import (
	"simplex/ctx"
	"simplex/node"
	"simplex/relate"
)

func BySideRelation(hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDirRelateValid(hull, cg)
}
