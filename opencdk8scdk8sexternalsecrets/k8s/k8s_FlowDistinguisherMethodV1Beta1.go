package k8s


// FlowDistinguisherMethod specifies the method of a flow distinguisher.
// Experimental.
type FlowDistinguisherMethodV1Beta1 struct {
	// `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace".
	//
	// Required.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

