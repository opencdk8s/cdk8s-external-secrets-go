package k8s


// ClusterRoleList is a collection of ClusterRoles.
// Experimental.
type KubeClusterRoleListProps struct {
	// Items is a list of ClusterRoles.
	// Experimental.
	Items *[]*KubeClusterRoleProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

