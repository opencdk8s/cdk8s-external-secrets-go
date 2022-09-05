package k8s


// A null or empty node selector term matches no objects.
//
// The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
// Experimental.
type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// Experimental.
	MatchExpressions *[]*NodeSelectorRequirement `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// A list of node selector requirements by node's fields.
	// Experimental.
	MatchFields *[]*NodeSelectorRequirement `field:"optional" json:"matchFields" yaml:"matchFields"`
}

