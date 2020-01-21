package constrain

import (
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/relate"
)

func ByGeometricRelation(hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsGeomRelateValid(hull, cg)
}
