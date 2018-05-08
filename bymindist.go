package constrain

import (
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/relate"
)

func ByMinDistRelation(options *opts.Opts, hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsDistRelateValid(options, hull, cg)
}
