package k8s


// Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
// Experimental.
type Preconditions struct {
	// Specifies the target ResourceVersion.
	// Experimental.
	ResourceVersion *string `field:"optional" json:"resourceVersion" yaml:"resourceVersion"`
	// Specifies the target UID.
	// Experimental.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

