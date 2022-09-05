package k8s


// RoleBindingList is a collection of RoleBindings.
// Experimental.
type KubeRoleBindingListProps struct {
	// Items is a list of RoleBindings.
	// Experimental.
	Items *[]*KubeRoleBindingProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

