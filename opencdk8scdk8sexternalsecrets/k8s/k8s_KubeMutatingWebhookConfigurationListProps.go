package k8s


// MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.
// Experimental.
type KubeMutatingWebhookConfigurationListProps struct {
	// List of MutatingWebhookConfiguration.
	// Experimental.
	Items *[]*KubeMutatingWebhookConfigurationProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

