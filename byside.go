package constrain

import (
	"simplex/ctx"
	"simplex/node"
	"simplex/relate"
	"simplex/lnr"
)

func BySideRelation(self lnr.Polygonal,hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDirRelateValid(self, hull, cg)
}
