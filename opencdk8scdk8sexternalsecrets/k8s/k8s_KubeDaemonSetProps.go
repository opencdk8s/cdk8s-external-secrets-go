package k8s


// DaemonSet represents the configuration of a daemon set.
// Experimental.
type KubeDaemonSetProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// The desired behavior of this daemon set.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// Experimental.
	Spec *DaemonSetSpec `field:"optional" json:"spec" yaml:"spec"`
}

