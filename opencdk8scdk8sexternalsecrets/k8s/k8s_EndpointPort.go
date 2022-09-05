package k8s


// EndpointPort is a tuple that describes a single port.
// Experimental.
type EndpointPort struct {
	// The port number of the endpoint.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The application protocol for this port.
	//
	// This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
	// Experimental.
	AppProtocol *string `field:"optional" json:"appProtocol" yaml:"appProtocol"`
	// The name of this port.
	//
	// This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IP protocol for this port.
	//
	// Must be UDP, TCP, or SCTP. Default is TCP.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

