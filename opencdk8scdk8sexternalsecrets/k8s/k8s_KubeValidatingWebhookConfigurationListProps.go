package k8s


// ValidatingWebhookConfigurationList is a list of ValidatingWebhookConfiguration.
// Experimental.
type KubeValidatingWebhookConfigurationListProps struct {
	// List of ValidatingWebhookConfiguration.
	// Experimental.
	Items *[]*KubeValidatingWebhookConfigurationProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

