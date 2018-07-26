package constrain

import (
    "github.com/TopoSimplify/opts"
    "github.com/TopoSimplify/node"
    "github.com/TopoSimplify/deform"
    "github.com/TopoSimplify/hdb"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func ByFeatureClassIntersection(options *opts.Opts, hull *node.Node, hulldb *hdb.Hdb, selections *[]*node.Node) bool {
    var bln = true
    //find hull neighbours
    var hulls = deform.SelectFeatureClass(options, hulldb, hull)
    for _, h := range hulls {
        //if bln & selection contains current hull : bln : false
        if bln && (h == hull) {
            bln = false // cmp ref
        }
        *selections = append(*selections, h)
    }
    return bln
}
