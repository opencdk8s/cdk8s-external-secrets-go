package k8s


// PodDNSConfigOption defines DNS resolver options of a pod.
// Experimental.
type PodDnsConfigOption struct {
	// Required.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

