// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretDataFromRemoteRef struct {
	// Experimental.
	Extract *ExternalSecretDataRemoteRef `field:"optional" json:"extract" yaml:"extract"`
	// Experimental.
	Find *ExternalSecretFind `field:"optional" json:"find" yaml:"find"`
	// Experimental.
	Rewrite *[]*ExternalSecretRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

