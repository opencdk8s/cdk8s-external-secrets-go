// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretDataRemoteRef struct {
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Experimental.
	ConversionStrategy *string `field:"optional" json:"conversionStrategy" yaml:"conversionStrategy"`
	// Experimental.
	DecodingStrategy *string `field:"optional" json:"decodingStrategy" yaml:"decodingStrategy"`
	// Experimental.
	MetadataPolicy *string `field:"optional" json:"metadataPolicy" yaml:"metadataPolicy"`
	// Experimental.
	Property *string `field:"optional" json:"property" yaml:"property"`
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

