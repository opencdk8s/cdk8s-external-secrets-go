package k8s


// EventSource contains information for an event.
// Experimental.
type EventSource struct {
	// Component from which the event is generated.
	// Experimental.
	Component *string `field:"optional" json:"component" yaml:"component"`
	// Node name on which the event is generated.
	// Experimental.
	Host *string `field:"optional" json:"host" yaml:"host"`
}

