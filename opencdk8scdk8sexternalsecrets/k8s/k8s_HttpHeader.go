package k8s


// HTTPHeader describes a custom header to be used in HTTP probes.
// Experimental.
type HttpHeader struct {
	// The header field name.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header field value.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

