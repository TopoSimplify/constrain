package constrain

import (
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/deform"
	"github.com/TopoSimplify/hdb"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func ByFeatureClassIntersection(
	options *opts.Opts,
	hull *node.Node,
	db *hdb.Hdb,
	selections *[]*node.Node,
) bool {
	var bln = true
	var hulls = deform.SelectFeatureClass(options, db, hull)
	for _, h := range hulls {
		if bln && (h == hull) {
			bln = false
		}
		*selections = append(*selections, h)
	}
	return bln
}
