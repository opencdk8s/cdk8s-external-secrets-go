// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type ExternalSecretData struct {
	// Experimental.
	RemoteRef *ExternalSecretDataRemoteRef `field:"required" json:"remoteRef" yaml:"remoteRef"`
	// Experimental.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

