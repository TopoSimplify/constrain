package constrain

import (
	"simplex/ctx"
	"simplex/node"
	"simplex/relate"
)

func BySideRelation(hull *node.Node, cgs *ctx.ContextGeometries) bool {
	return relate.IsDirRelateValid(hull, cgs)
}
