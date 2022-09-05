package k8s


// AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole.
// Experimental.
type AggregationRule struct {
	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules.
	//
	// If any of the selectors match, then the ClusterRole's permissions will be added.
	// Experimental.
	ClusterRoleSelectors *[]*LabelSelector `field:"optional" json:"clusterRoleSelectors" yaml:"clusterRoleSelectors"`
}

