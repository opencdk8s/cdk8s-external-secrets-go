package k8s


// A node selector represents the union of the results of one or more label queries over a set of nodes;
//
// that is, it represents the OR of the selectors represented by the node selector terms.
// Experimental.
type NodeSelector struct {
	// Required.
	//
	// A list of node selector terms. The terms are ORed.
	// Experimental.
	NodeSelectorTerms *[]*NodeSelectorTerm `field:"required" json:"nodeSelectorTerms" yaml:"nodeSelectorTerms"`
}

