package k8s


// Binding ties one object to another;
//
// for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.
// Experimental.
type KubeBindingProps struct {
	// The target object that you want to bind to the standard object.
	// Experimental.
	Target *ObjectReference `field:"required" json:"target" yaml:"target"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

