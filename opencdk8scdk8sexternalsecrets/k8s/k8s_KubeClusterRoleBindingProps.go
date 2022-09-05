package k8s


// ClusterRoleBinding references a ClusterRole, but not contain it.
//
// It can reference a ClusterRole in the global namespace, and adds who information via Subject.
// Experimental.
type KubeClusterRoleBindingProps struct {
	// RoleRef can only reference a ClusterRole in the global namespace.
	//
	// If the RoleRef cannot be resolved, the Authorizer must return an error.
	// Experimental.
	RoleRef *RoleRef `field:"required" json:"roleRef" yaml:"roleRef"`
	// Standard object's metadata.
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Subjects holds references to the objects the role applies to.
	// Experimental.
	Subjects *[]*Subject `field:"optional" json:"subjects" yaml:"subjects"`
}

