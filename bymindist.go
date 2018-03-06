package constrain

import (
	"simplex/opts"
	"simplex/node"
	"simplex/ctx"
	"simplex/relate"
)

func ByMinDistRelation(options *opts.Opts, hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsDistRelateValid(options, hull, cg)
}
