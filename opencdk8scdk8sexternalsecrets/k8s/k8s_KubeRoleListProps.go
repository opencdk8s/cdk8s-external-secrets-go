package k8s


// RoleList is a collection of Roles.
// Experimental.
type KubeRoleListProps struct {
	// Items is a list of Roles.
	// Experimental.
	Items *[]*KubeRoleProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

