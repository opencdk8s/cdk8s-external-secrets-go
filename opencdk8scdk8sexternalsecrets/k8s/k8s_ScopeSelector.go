package k8s


// A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
// Experimental.
type ScopeSelector struct {
	// A list of scope selector requirements by scope of the resources.
	// Experimental.
	MatchExpressions *[]*ScopedResourceSelectorRequirement `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
}

