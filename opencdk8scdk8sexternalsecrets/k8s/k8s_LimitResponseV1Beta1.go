package k8s


// LimitResponse defines how to handle requests that can not be executed right now.
// Experimental.
type LimitResponseV1Beta1 struct {
	// `type` is "Queue" or "Reject".
	//
	// "Queue" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. "Reject" means that requests that can not be executed upon arrival are rejected. Required.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `queuing` holds the configuration parameters for queuing.
	//
	// This field may be non-empty only if `type` is `"Queue"`.
	// Experimental.
	Queuing *QueuingConfigurationV1Beta1 `field:"optional" json:"queuing" yaml:"queuing"`
}

