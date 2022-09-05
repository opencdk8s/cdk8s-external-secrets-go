package k8s


// ClusterRoleBindingList is a collection of ClusterRoleBindings.
//
// Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBindings, and will no longer be served in v1.22.
// Experimental.
type KubeClusterRoleBindingListV1Alpha1Props struct {
	// Items is a list of ClusterRoleBindings.
	// Experimental.
	Items *[]*KubeClusterRoleBindingV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

