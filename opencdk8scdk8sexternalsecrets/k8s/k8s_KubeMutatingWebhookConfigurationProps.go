package k8s


// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object.
// Experimental.
type KubeMutatingWebhookConfigurationProps struct {
	// Standard object metadata;
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	// Experimental.
	Webhooks *[]*MutatingWebhook `field:"optional" json:"webhooks" yaml:"webhooks"`
}

