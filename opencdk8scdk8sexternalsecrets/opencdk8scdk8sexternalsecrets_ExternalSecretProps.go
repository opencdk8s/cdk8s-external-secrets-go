// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets

import (
	"github.com/opencdk8s/cdk8s-external-secrets-go/opencdk8scdk8sexternalsecrets/k8s"
)

// Experimental.
type ExternalSecretProps struct {
	// Experimental.
	Metadata *k8s.ObjectMeta `field:"required" json:"metadata" yaml:"metadata"`
	// Experimental.
	Spec *ExternalSecretSpec `field:"required" json:"spec" yaml:"spec"`
}

