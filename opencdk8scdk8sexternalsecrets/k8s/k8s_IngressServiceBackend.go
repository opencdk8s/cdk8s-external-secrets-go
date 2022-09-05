package k8s


// IngressServiceBackend references a Kubernetes Service as a Backend.
// Experimental.
type IngressServiceBackend struct {
	// Name is the referenced service.
	//
	// The service must exist in the same namespace as the Ingress object.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port of the referenced service.
	//
	// A port name or port number is required for a IngressServiceBackend.
	// Experimental.
	Port *ServiceBackendPort `field:"optional" json:"port" yaml:"port"`
}

