// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretFind struct {
	// Experimental.
	ConversionStrategy *string `field:"optional" json:"conversionStrategy" yaml:"conversionStrategy"`
	// Experimental.
	DecodingStrategy *string `field:"optional" json:"decodingStrategy" yaml:"decodingStrategy"`
	// Experimental.
	Name *FindName `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

