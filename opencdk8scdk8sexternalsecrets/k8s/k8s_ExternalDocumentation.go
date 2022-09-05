package k8s


// ExternalDocumentation allows referencing an external resource for extended documentation.
// Experimental.
type ExternalDocumentation struct {
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

