package constrain

import (
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/node"
	"relate"
)

func BySideRelation(hull *node.Node, cgs *ctx.ContextGeometries) bool {
	return relate.IsDirRelateValid(hull, cgs)
}
