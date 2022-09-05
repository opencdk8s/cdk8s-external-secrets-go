package k8s


// Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
// Experimental.
type KubeRoleProps struct {
	// Standard object's metadata.
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Rules holds all the PolicyRules for this Role.
	// Experimental.
	Rules *[]*PolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

