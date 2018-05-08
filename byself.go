package constrain

import (
    "github.com/intdxdt/rtree"
    "github.com/TopoSimplify/node"
    "github.com/TopoSimplify/opts"
    "github.com/TopoSimplify/deform"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func BySelfIntersection(options *opts.Opts, hull *node.Node, hulldb *rtree.RTree, selections *[]*node.Node) bool {
    //assume hull is valid and proof otherwise
    var bln = true
    // find hull neighbours
    var hulls = deform.Select(options, hulldb, hull)
    for _, h := range hulls {
        //if bln & selection contains current hull : bln : false
        if bln && (h == hull) {
            bln = false //cmp &
        }
        *selections = append(*selections, h)
    }

    return bln
}
