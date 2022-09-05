package k8s


// CustomResourceConversion describes how to convert different versions of a CR.
// Experimental.
type CustomResourceConversion struct {
	// strategy specifies how custom resources are converted between versions.
	//
	// Allowed values are: - `None`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `Webhook`: API Server will call to an external webhook to do the conversion. Additional information
	// is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
	// Experimental.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// webhook describes how to call the conversion webhook.
	//
	// Required when `strategy` is set to `Webhook`.
	// Experimental.
	Webhook *WebhookConversion `field:"optional" json:"webhook" yaml:"webhook"`
}

