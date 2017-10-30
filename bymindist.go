package constrain

import (
	"simplex/opts"
	"simplex/node"
	"simplex/ctx"
	"simplex/relate"
)

func ByMinDistRelation(options *opts.Opts, hull *node.Node, cg *ctx.ContextGeometry) bool {
	return relate.IsDistRelateValid(options, hull, cg)
}
