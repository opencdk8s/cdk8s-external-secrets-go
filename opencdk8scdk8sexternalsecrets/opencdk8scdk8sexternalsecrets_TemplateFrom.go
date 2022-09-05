// @opencdk8s/cdk8s-external-secrets
package opencdk8scdk8sexternalsecrets


// Experimental.
type TemplateFrom struct {
	// Experimental.
	ConfigMap *TemplateRef `field:"optional" json:"configMap" yaml:"configMap"`
	// Experimental.
	Secret *TemplateRef `field:"optional" json:"secret" yaml:"secret"`
}

