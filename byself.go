package constrain

import (
	"github.com/TopoSimplify/hdb"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/cmap"
	"github.com/TopoSimplify/deform"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func BySelfIntersection(options *opts.Opts, hull *node.Node,
	hulldb *hdb.Hdb, selections *[]*node.Node, cache *cmap.CacheMap) bool {
	//assume hull is valid and proof otherwise
	var bln = true
	// find hull neighbours
	var hulls = deform.Select(options, hulldb, hull, cache)
	for _, h := range hulls {
		//if bln & selection contains current hull : bln : false
		if bln && (h == hull) {
			bln = false //cmp &
		}
		*selections = append(*selections, h)
	}

	return bln
}
