package k8s


// Overhead structure represents the resource overhead associated with running a pod.
// Experimental.
type OverheadV1Beta1 struct {
	// PodFixed represents the fixed resource overhead associated with running a pod.
	// Experimental.
	PodFixed *map[string]Quantity `field:"optional" json:"podFixed" yaml:"podFixed"`
}

