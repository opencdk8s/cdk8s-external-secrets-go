package k8s


// ClusterRoleList is a collection of ClusterRoles.
//
// Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoles, and will no longer be served in v1.22.
// Experimental.
type KubeClusterRoleListV1Alpha1Props struct {
	// Items is a list of ClusterRoles.
	// Experimental.
	Items *[]*KubeClusterRoleV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

