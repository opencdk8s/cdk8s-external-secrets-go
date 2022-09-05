package k8s


// ClusterRoleBindingList is a collection of ClusterRoleBindings.
// Experimental.
type KubeClusterRoleBindingListProps struct {
	// Items is a list of ClusterRoleBindings.
	// Experimental.
	Items *[]*KubeClusterRoleBindingProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

