// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretSpec struct {
	// Experimental.
	SecretStoreRef *SecretStoreRef `field:"required" json:"secretStoreRef" yaml:"secretStoreRef"`
	// Experimental.
	Data *[]*ExternalSecretData `field:"optional" json:"data" yaml:"data"`
	// Experimental.
	DataFrom *[]*ExternalSecretDataFromRemoteRef `field:"optional" json:"dataFrom" yaml:"dataFrom"`
	// Experimental.
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// Experimental.
	Target *ExternalSecretTarget `field:"optional" json:"target" yaml:"target"`
}

