package k8s


// Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
// Experimental.
type Scheduling struct {
	// nodeSelector lists labels that must be present on nodes that support this RuntimeClass.
	//
	// Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
	// Experimental.
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
	// Experimental.
	Tolerations *[]*Toleration `field:"optional" json:"tolerations" yaml:"tolerations"`
}

