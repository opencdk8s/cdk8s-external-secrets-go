package k8s


// RunAsUserStrategyOptions defines the strategy type and any options used to create the strategy.
// Experimental.
type RunAsUserStrategyOptionsV1Beta1 struct {
	// rule is the strategy that will dictate the allowable RunAsUser values that may be set.
	// Experimental.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// ranges are the allowed ranges of uids that may be used.
	//
	// If you would like to force a single uid then supply a single range with the same start and end. Required for MustRunAs.
	// Experimental.
	Ranges *[]*IdRangeV1Beta1 `field:"optional" json:"ranges" yaml:"ranges"`
}

