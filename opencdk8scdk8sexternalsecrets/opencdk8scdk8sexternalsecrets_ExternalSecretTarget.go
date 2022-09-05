// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretTarget struct {
	// Experimental.
	CreationPolicy *string `field:"optional" json:"creationPolicy" yaml:"creationPolicy"`
	// Experimental.
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Experimental.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Template *ExternalSecretTemplate `field:"optional" json:"template" yaml:"template"`
}

