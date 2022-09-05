package k8s


// RoleRef contains information that points to the role being used.
// Experimental.
type RoleRef struct {
	// APIGroup is the group for the resource being referenced.
	// Experimental.
	ApiGroup *string `field:"required" json:"apiGroup" yaml:"apiGroup"`
	// Kind is the type of resource being referenced.
	// Experimental.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name is the name of resource being referenced.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

