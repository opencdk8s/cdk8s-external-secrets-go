// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretTemplate struct {
	// Experimental.
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Experimental.
	Metadata *ExternalSecretTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Experimental.
	TemplateFrom *[]*TemplateFrom `field:"optional" json:"templateFrom" yaml:"templateFrom"`
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

