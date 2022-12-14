package k8s


// RoleBindingList is a collection of RoleBindings Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleBindingList, and will no longer be served in v1.22.
// Experimental.
type KubeRoleBindingListV1Alpha1Props struct {
	// Items is a list of RoleBindings.
	// Experimental.
	Items *[]*KubeRoleBindingV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

