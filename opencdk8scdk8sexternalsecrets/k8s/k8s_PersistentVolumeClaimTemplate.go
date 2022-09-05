package k8s


// PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
// Experimental.
type PersistentVolumeClaimTemplate struct {
	// The specification for the PersistentVolumeClaim.
	//
	// The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.
	// Experimental.
	Spec *PersistentVolumeClaimSpec `field:"required" json:"spec" yaml:"spec"`
	// May contain labels and annotations that will be copied into the PVC when creating it.
	//
	// No other fields are allowed and will be rejected during validation.
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

