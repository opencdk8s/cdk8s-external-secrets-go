package k8s


// SELinuxOptions are the labels to be applied to the container.
// Experimental.
type SeLinuxOptions struct {
	// Level is SELinux level label that applies to the container.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Role is a SELinux role label that applies to the container.
	// Experimental.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Type is a SELinux type label that applies to the container.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User is a SELinux user label that applies to the container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
}

